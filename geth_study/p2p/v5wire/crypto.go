// Copyright 2020 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package v5wire

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/ecdsa"
	"crypto/elliptic"
	"errors"
	"fmt"
	"hash"

	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/p2p/enode"
	"golang.org/x/crypto/hkdf"
)

const (
	// Encryption/authentication parameters.
	// 암호화/인증 매개 변수입니다.
	aesKeySize   = 16
	gcmNonceSize = 12
)

// Nonce represents a nonce used for AES/GCM.
// 논스는 AES/GCM에 사용되는 논스를 나타냅니다.
type Nonce [gcmNonceSize]byte

// EncodePubkey encodes a public key.
// EncodPubkey는 공개 키를 인코딩합니다.
func EncodePubkey(key *ecdsa.PublicKey) []byte {
	switch key.Curve {
	case crypto.S256():
		return crypto.CompressPubkey(key)
	default:
		panic("unsupported curve " + key.Curve.Params().Name + " in EncodePubkey")
	}
}

// DecodePubkey decodes a public key in compressed format.
// DecodePubkey는 공개 키를 압축된 형식으로 디코딩합니다.
func DecodePubkey(curve elliptic.Curve, e []byte) (*ecdsa.PublicKey, error) {
	switch curve {
	case crypto.S256():
		if len(e) != 33 {
			return nil, errors.New("wrong size public key data")
		}
		return crypto.DecompressPubkey(e)
	default:
		return nil, fmt.Errorf("unsupported curve %s in DecodePubkey", curve.Params().Name)
	}
}

// idNonceHash computes the ID signature hash used in the handshake.
// idNonceHash는 핸드셰이크에 사용되는 ID 서명 해시를 계산합니다.
func idNonceHash(h hash.Hash, challenge, ephkey []byte, destID enode.ID) []byte {
	h.Reset()
	h.Write([]byte("discovery v5 identity proof"))
	h.Write(challenge)
	h.Write(ephkey)
	h.Write(destID[:])
	return h.Sum(nil)
}

// makeIDSignature creates the ID nonce signature.
// makeIDsignature는 ID nonce 시그니처를 만듭니다.
func makeIDSignature(hash hash.Hash, key *ecdsa.PrivateKey, challenge, ephkey []byte, destID enode.ID) ([]byte, error) {
	input := idNonceHash(hash, challenge, ephkey, destID)
	switch key.Curve {
	case crypto.S256():
		idsig, err := crypto.Sign(input, key)
		if err != nil {
			return nil, err
		}
		return idsig[:len(idsig)-1], nil // remove recovery ID 복구 ID 제거
	default:
		return nil, fmt.Errorf("unsupported curve %s", key.Curve.Params().Name)
	}
}

// s256raw is an unparsed secp256k1 public key ENR entry.
// s256raw는 구문 분석되지 않은 secp256k1 공개 키 ENR 항목입니다.
type s256raw []byte

func (s256raw) ENRKey() string { return "secp256k1" }

// verifyIDSignature checks that signature over idnonce was made by the given node.
// ID 서명이 지정된 노드에 의해 이루어졌는지 확인합니다.
func verifyIDSignature(hash hash.Hash, sig []byte, n *enode.Node, challenge, ephkey []byte, destID enode.ID) error {
	switch idscheme := n.Record().IdentityScheme(); idscheme {
	case "v4":
		var pubkey s256raw
		if n.Load(&pubkey) != nil {
			return errors.New("no secp256k1 public key in record")
		}
		input := idNonceHash(hash, challenge, ephkey, destID)
		if !crypto.VerifySignature(pubkey, input, sig) {
			return errInvalidNonceSig
		}
		return nil
	default:
		return fmt.Errorf("can't verify ID nonce signature against scheme %q", idscheme)
	}
}

type hashFn func() hash.Hash

// deriveKeys creates the session keys.
// derimKeys는 세션 키를 생성합니다.
func deriveKeys(hash hashFn, priv *ecdsa.PrivateKey, pub *ecdsa.PublicKey, n1, n2 enode.ID, challenge []byte) *session {
	const text = "discovery v5 key agreement"
	var info = make([]byte, 0, len(text)+len(n1)+len(n2))
	info = append(info, text...)
	info = append(info, n1[:]...)
	info = append(info, n2[:]...)

	eph := ecdh(priv, pub)
	if eph == nil {
		return nil
	}
	kdf := hkdf.New(hash, eph, challenge, info)
	sec := session{writeKey: make([]byte, aesKeySize), readKey: make([]byte, aesKeySize)}
	kdf.Read(sec.writeKey)
	kdf.Read(sec.readKey)
	for i := range eph {
		eph[i] = 0
	}
	return &sec
}

// ecdh creates a shared secret.
// ecdh는 공유 비밀을 만듭니다.
func ecdh(privkey *ecdsa.PrivateKey, pubkey *ecdsa.PublicKey) []byte {
	secX, secY := pubkey.ScalarMult(pubkey.X, pubkey.Y, privkey.D.Bytes())
	if secX == nil {
		return nil
	}
	sec := make([]byte, 33)
	sec[0] = 0x02 | byte(secY.Bit(0))
	math.ReadBits(secX, sec[1:])
	return sec
}

// encryptGCM encrypts pt using AES-GCM with the given key and nonce. The ciphertext is
// appended to dest, which must not overlap with plaintext. The resulting ciphertext is 16
// bytes longer than plaintext because it contains an authentication tag.
// 암호화하다GCM은 지정된 키와 난스를 사용하여 AES-GCM을 사용하여 pt를 암호화합니다. 암호문은
// 일반 텍스트와 겹치지 않아야 하는 대상에 추가됩니다. 결과 암호문은 16이다
// 인증 태그를 포함하므로 일반 텍스트보다 바이트가 깁니다.
func encryptGCM(dest, key, nonce, plaintext, authData []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(fmt.Errorf("can't create block cipher: %v", err))
	}
	aesgcm, err := cipher.NewGCMWithNonceSize(block, gcmNonceSize)
	if err != nil {
		panic(fmt.Errorf("can't create GCM: %v", err))
	}
	return aesgcm.Seal(dest, nonce, plaintext, authData), nil
}

// decryptGCM decrypts ct using AES-GCM with the given key and nonce.
// decryptGCM은 AES-GCM을 사용하여 지정된 키와 난스를 사용하여 ct를 해독합니다.
func decryptGCM(key, nonce, ct, authData []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("can't create block cipher: %v", err)
	}
	if len(nonce) != gcmNonceSize {
		return nil, fmt.Errorf("invalid GCM nonce size: %d", len(nonce))
	}
	aesgcm, err := cipher.NewGCMWithNonceSize(block, gcmNonceSize)
	if err != nil {
		return nil, fmt.Errorf("can't create GCM: %v", err)
	}
	pt := make([]byte, 0, len(ct))
	return aesgcm.Open(pt, nonce, ct, authData)
}
