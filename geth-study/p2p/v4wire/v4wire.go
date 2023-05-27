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

// Package v4wire implements the Discovery v4 Wire Protocol.
// Copyright 2020 The Go-Ethereum 권위자들
// 이 파일은 Go-etherum 라이브러리의 일부입니다.
//
// Go-etherum 라이브러리는 무료 소프트웨어입니다. 재배포 및/또는 수정할 수 있습니다
// 그것은 GNU 소규모 일반 공중 사용 허가서의 조건에 의해 출판되었다
// Free Software Foundation, License 버전 3 또는
// (선택사항에 따라) 이후 버전을 선택합니다.
//
// Go-etherum 라이브러리는 유용하기를 바라며 배포된다,
// 하지만 어떠한 보증도 없이, 심지어 묵시적인 보증도 없다
// 특정 목적에 대한 상품성 또는 적합성. 을 참조
// 자세한 내용은 GNU Lesser General Public License를 참조하십시오.
//
// 당신은 GNU 소규모 일반 공중 사용 허가서 사본을 받았어야 했다
// 라이브러리와 함께. 그렇지 않은 경우 <http://www.gnu.org/licenses/>을 참조하십시오.

// 패키지 v4wire는 Discovery v4 Wire 프로토콜을 구현합니다.
package v4wire

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"errors"
	"fmt"
	"math/big"
	"net"
	"time"

	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/p2p/enode"
	"github.com/ethereum/go-ethereum/p2p/enr"
	"github.com/ethereum/go-ethereum/rlp"
)

// RPC packet types
// RPC 패킷 타입
const (
	PingPacket = iota + 1 // zero is 'reserved'
	PongPacket
	FindnodePacket
	NeighborsPacket
	ENRRequestPacket
	ENRResponsePacket
)

// RPC request structures
// RPC 요청 구조
/*
RLP stands for Recursive Length Prefix, it is a serialization/deserialization encoding scheme used in the go-ethereum codebase.
It is used to serialize data structures for storage on the blockchain and for transmitting data between nodes.
RLP is a binary encoding that recursively encodes data structures with variable-length elements.
It is particularly useful for encoding lists, as it can represent the length of the list in a compact and efficient manner.
In Ethereum, RLP is used to encode transactions, blocks, and account state data, among other things.

RLP는 Recursive Length Prefix의 약자로, Geth 코드베이스에 사용되는 직렬화/비직렬화 인코딩 방식이다.
블록체인에 저장하기 위한 데이터 구조를 직렬화하고 노드 간 데이터를 전송하는 데 사용된다.
RLP는 가변 길이 요소로 데이터 구조를 재귀적으로 인코딩하는 이진 인코딩이다.
목록을 인코딩하는 데 특히 유용한데, 목록의 길이를 간결하고 효율적으로 나타낼 수 있기 때문이다.
이더리움에서 RLP는 트랜잭션, 블록, 계정 상태 데이터 등을 인코딩하는 데 사용된다.



Recursive Length Prefix (RLP) serialization은 Ethereum의 실행 클라이언트에서 광범위하게 사용됩니다.
RLP는 노드 간 데이터 전송을 공간 효율적인 형식으로 표준화합니다.
RLP의 목적은 이진 데이터의 임의로 중첩된 배열을 인코딩하는 것이며,
RLP는 Ethereum의 실행 레이어에서 객체를 직렬화하는 데 사용되는 기본 인코딩 방법입니다.
RLP의 목적은 구조를 인코딩하는 것 뿐이며, 특정 데이터 유형 (예: 문자열, 부동 소수점)의 인코딩은 상위 프로토콜에서 결정됩니다.
하지만 양의 RLP 정수는 선행하는 0이 없는 빅 엔디언 이진 형식으로 표시되어야 합니다 (따라서 정수 값 0은 빈 바이트 배열과 같습니다).
선행하는 0이 있는 역직렬화된 양의 정수는 잘못 처리됩니다. 문자열 길이의 정수 표현도 이와 같이 인코딩되어야 하며, 페이로드의 정수도 동일합니다.

RLP를 사용하여 사전을 인코딩하려면, 두 가지 권장되는 표준형이 있습니다:

lexicographic 순서의 키를 사용하여 [[k1, v1], [k2, v2]...]를 사용합니다.
Ethereum과 같은 상위 수준 Patricia Tree 인코딩을 사용합니다.


*/
type (
	Ping struct {
		Version    uint
		From, To   Endpoint
		Expiration uint64
		ENRSeq     uint64 `rlp:"optional"` // Sequence number of local record, added by EIP-868. EIP-868에 의해 추가된 로컬 레코드의 시퀀스 번호.

		// Ignore additional fields (for forward compatibility).
		// 앞에있는 호환성을 위해 추가 필드를 무시합니다.
		Rest []rlp.RawValue `rlp:"tail"`
	}

	// Pong is the reply to ping.
	// 퐁은 핑에 대한 답이다.
	Pong struct {
		// This field should mirror the UDP envelope address
		// of the ping packet, which provides a way to discover the
		// external address (after NAT).
		// 이 필드는 UDP 엔벨로프 주소를 미러링해야 합니다
		// ping 패킷을 검색하는 방법을 제공합니다
		// 외부 주소(NAT 이후).
		To         Endpoint
		ReplyTok   []byte // This contains the hash of the ping packet. 이것은 ping 패킷의 해시를 포함합니다.
		Expiration uint64 // Absolute timestamp at which the packet becomes invalid. 패킷이 유효하지 않게 되는 절대 타임스탬프입니다.
		ENRSeq     uint64 `rlp:"optional"` // Sequence number of local record, added by EIP-868. EIP-868에 의해 추가된 로컬 레코드의 시퀀스 번호.

		// Ignore additional fields (for forward compatibility).
		// 앞에있는 호환성을 위해 추가 필드를 무시합니다.
		Rest []rlp.RawValue `rlp:"tail"`
	}

	// Findnode is a query for nodes close to the given target.
	// 노드 찾기는 지정된 대상에 가까운 노드에 대한 쿼리입니다
	Findnode struct {
		Target     Pubkey
		Expiration uint64
		// Ignore additional fields (for forward compatibility).
		// 앞에있는 호환성을 위해 추가 필드를 무시합니다.
		Rest []rlp.RawValue `rlp:"tail"`
	}

	// Neighbors is the reply to findnode.
	// Neighbors는 노드 찾기에 대한 응답입니다.
	Neighbors struct {
		Nodes      []Node
		Expiration uint64
		// Ignore additional fields (for forward compatibility).
		// 앞에있는 호환성을 위해 추가 필드를 무시합니다.
		Rest []rlp.RawValue `rlp:"tail"`
	}
	/*
	   ENR stands for Ethereum Name Service Record, which is a type of record used for identifying and discovering nodes in the Ethereum network.
	   ENR requests in geth refer to requests made by nodes to retrieve the ENR records of other nodes.

	   When a node joins the Ethereum network, it creates its own ENR record, which contains information about its network address,
	   public key, and other metadata. Other nodes can then use this record to discover and connect to the node.

	   ENR requests are sent when a node needs to connect to another node for various reasons, such as syncing its blockchain or sending a transaction. T
	   he node sends an ENR request to the target node, which responds with its own ENR record.
	   This allows the requesting node to obtain the necessary information to establish a connection with the target node.

	   ENR은 이더리움 네임 서비스 레코드(Ethereum Name Service Record)의 약자로, 이더리움 네트워크에서 노드를 식별하고 발견하는 데 사용되는 레코드의 일종이다.
	   ENR 요청은 노드가 다른 노드의 ENR 레코드를 검색하기 위해 수행한 요청을 의미합니다.

	   노드가 이더리움 네트워크에 가입하면 네트워크 주소, 공개 키 및 기타 메타데이터에 대한 정보를 포함하는 자체 ENR 레코드를 생성한다.
	   그러면 다른 노드가 이 레코드를 사용하여 노드를 검색하고 연결할 수 있습니다.

	   ENR 요청은 블록체인을 동기화하거나 트랜잭션을 전송하는 등 다양한 이유로 노드가 다른 노드에 연결해야 할 때 전송됩니다.
	   노드는 대상 노드에 ENR 요청을 전송하고, 대상 노드는 자체 ENR 레코드로 응답합니다. 이를 통해 요청 노드는 대상 노드와의 연결을 설정하는 데 필요한 정보를 얻을 수 있습니다
	*/

	// ENRRequest queries for the remote node's record.
	// 원격 노드의 레코드에 대한 ENR 요청 쿼리입니다.
	ENRRequest struct {
		Expiration uint64
		// Ignore additional fields (for forward compatibility).
		Rest []rlp.RawValue `rlp:"tail"`
	}

	// ENRResponse is the reply to ENRRequest.
	// ENR 응답은 ENR 요청에 대한 응답입니다.
	ENRResponse struct {
		ReplyTok []byte // Hash of the ENRRequest packet. ENRRequest 패킷의 해시입니다.
		Record   enr.Record
		// Ignore additional fields (for forward compatibility).
		Rest []rlp.RawValue `rlp:"tail"`
	}
)

// MaxNeighbors is the maximum number of neighbor nodes in a Neighbors packet.
// MaxNeighbors는 Neighbors 패킷의 최대 인접 노드 수입니다.
const MaxNeighbors = 12

// This code computes the MaxNeighbors constant value.

// func init() {
// 	var maxNeighbors int
// 	p := Neighbors{Expiration: ^uint64(0)}
// 	maxSizeNode := Node{IP: make(net.IP, 16), UDP: ^uint16(0), TCP: ^uint16(0)}
// 	for n := 0; ; n++ {
// 		p.Nodes = append(p.Nodes, maxSizeNode)
// 		size, _, err := rlp.EncodeToReader(p)
// 		if err != nil {
// 			// If this ever happens, it will be caught by the unit tests.
// 			panic("cannot encode: " + err.Error())
// 		}
// 		if headSize+size+1 >= 1280 {
// 			maxNeighbors = n
// 			break
// 		}
// 	}
// 	fmt.Println("maxNeighbors", maxNeighbors)
// }

// Pubkey represents an encoded 64-byte secp256k1 public key.
// Pubkey는 인코딩된 64바이트 secp256k1 공개 키를 나타냅니다.
type Pubkey [64]byte

// ID returns the node ID corresponding to the public key.
func (e Pubkey) ID() enode.ID {
	return enode.ID(crypto.Keccak256Hash(e[:]))
}

// Node represents information about a node.
// 노드는 노드에 대한 정보를 나타냅니다.
type Node struct {
	IP  net.IP // len 4 for IPv4 or 16 for IPv6 IPv4의 경우 len 4 또는 IPv6의 경우 16
	UDP uint16 // for discovery protocol
	// 사용자 데이터그램 프로토콜(User Datagram Protocol, UDP)
	// UDP의 전송 방식은 너무 단순해서 서비스의 신뢰성이 낮고, 데이터그램 도착 순서가 바뀌거나, 중복되거나, 심지어는 통보 없이 누락시키기도 한다.
	// UDP는 일반적으로 오류의 검사와 수정이 필요 없는 애플리케이션에서 수행할 것으로 가정한다.
	// TCP는 데이터를 주고 받을 양단 간에 먼저 연결을 설정하고 설정된 연결을 통해 양방향으로 데이터를 전송하지만, UDP는 연결을 설정하지 않고 수신자가 데이터를 받을 준비를 확인하는 단계를 거치지 않고 단방향으로 정보를 전송한다.
	// 신뢰성 - TCP는 메시지 수신을 확인하지만 UDP는 수신자가 메시지를 수신했는지 확인할 수 없다.
	// 순서 정렬 - TCP에서는 메시지가 보내진 순서를 보장하기 위해 재조립하지만 UDP는 메시지 도착 순서를 예측할 수 없다.
	// 부하 - TCP보다 속도가 일반적으로 빠르고 오버헤드가 적다.

	TCP uint16 // for RLPx protocol
	// 전송 제어 프로토콜(Transmission Control Protocol, TCP, 문화어: 전송조종규약)은 인터넷 프로토콜 스위트(IP)의 핵심 프로토콜 중 하나로,
	// IP와 함께 TCP/IP라는 명칭으로도 널리 불린다. TCP는 근거리 통신망이나 인트라넷, 인터넷에 연결된 컴퓨터에서 실행되는 프로그램 간에 일련의 옥텟을 안정적으로,
	// 순서대로, 에러없이 교환할 수 있게 한다. TCP는 전송 계층에 위치한다.
	// 네트워크의 정보 전달을 통제하는 프로토콜이자 인터넷을 이루는 핵심 프로토콜의 하나로서 국제 인터넷 표준화 기구(IETF)의 RFC 793에 기술되어 있다.
	ID Pubkey
}

/*


 */

// Endpoint represents a network endpoint.
// 엔드포인트는 네트워크 엔드포인트를 나타냅니다.
type Endpoint struct {
	IP  net.IP // len 4 for IPv4 or 16 for IPv6
	UDP uint16 // for discovery protocol
	TCP uint16 // for RLPx protocol
}

// NewEndpoint creates an endpoint.
// 새 엔드포인트는 엔드포인트를 만듭니다.
func NewEndpoint(addr *net.UDPAddr, tcpPort uint16) Endpoint {
	ip := net.IP{}
	if ip4 := addr.IP.To4(); ip4 != nil {
		ip = ip4
	} else if ip6 := addr.IP.To16(); ip6 != nil {
		ip = ip6
	}
	fmt.Println("NewEndpoint")
	return Endpoint{IP: ip, UDP: uint16(addr.Port), TCP: tcpPort}
}

type Packet interface {
	// Name is the name of the package, for logging purposes.
	// Name은 로깅을 위한 패키지의 이름입니다.
	Name() string
	// Kind is the packet type, for logging purposes.
	// Kind는 로깅을 위한 패킷 유형입니다.
	Kind() byte
}

// ping 함수
func (req *Ping) Name() string { return "PING/v4" }
func (req *Ping) Kind() byte   { return PingPacket }

// pong 함수
func (req *Pong) Name() string { return "PONG/v4" }
func (req *Pong) Kind() byte   { return PongPacket }

// Findnode 함수 지정된 대상에 가까운 노드에 대한 쿼리입니다
func (req *Findnode) Name() string { return "FINDNODE/v4" }
func (req *Findnode) Kind() byte   { return FindnodePacket }

// Findnode 함수에 대한 응답
func (req *Neighbors) Name() string { return "NEIGHBORS/v4" }
func (req *Neighbors) Kind() byte   { return NeighborsPacket }

// Ethereum Name Service Record 요청은 노드가 다른 노드의 ENR 레코드를 검색하기 위해 수행한 요청을 의미합니다.
func (req *ENRRequest) Name() string { return "ENRREQUEST/v4" }
func (req *ENRRequest) Kind() byte   { return ENRRequestPacket }

// Ethereum Name Service Record 응답은  자체 ENR 레코드로 응답합니다.  (노드가 이더리움 네트워크에 가입하면 네트워크 주소, 공개 키 및 기타 메타데이터에 대한 정보를 포함하는 자체 ENR 레코드를 생성한다.)
func (req *ENRResponse) Name() string { return "ENRRESPONSE/v4" }
func (req *ENRResponse) Kind() byte   { return ENRResponsePacket }

// Expired checks whether the given UNIX time stamp is in the past.
// Expired는 주어진 UNIX 시간 스탬프가 과거인지 확인합니다.
func Expired(ts uint64) bool {
	return time.Unix(int64(ts), 0).Before(time.Now())
}

// Encoder/decoder.
// 엔코드/디코드
const (
	macSize  = 32
	sigSize  = crypto.SignatureLength
	headSize = macSize + sigSize // space of packet frame data
)

var (
	ErrPacketTooSmall = errors.New("too small")
	ErrBadHash        = errors.New("bad hash")
	ErrBadPoint       = errors.New("invalid curve point")
)

var headSpace = make([]byte, headSize)

// Decode reads a discovery v4 packet.
// Decode은 discovery v4 패킷을 읽습니다.
func Decode(input []byte) (Packet, Pubkey, []byte, error) {
	if len(input) < headSize+1 {
		return nil, Pubkey{}, nil, ErrPacketTooSmall
	}
	hash, sig, sigdata := input[:macSize], input[macSize:headSize], input[headSize:]
	shouldhash := crypto.Keccak256(input[macSize:])
	if !bytes.Equal(hash, shouldhash) {
		return nil, Pubkey{}, nil, ErrBadHash
	}
	fromKey, err := recoverNodeKey(crypto.Keccak256(input[headSize:]), sig)
	if err != nil {
		return nil, fromKey, hash, err
	}

	var req Packet
	switch ptype := sigdata[0]; ptype {
	case PingPacket:
		req = new(Ping)
	case PongPacket:
		req = new(Pong)
	case FindnodePacket:
		req = new(Findnode)
	case NeighborsPacket:
		req = new(Neighbors)
	case ENRRequestPacket:
		req = new(ENRRequest)
	case ENRResponsePacket:
		req = new(ENRResponse)
	default:
		return nil, fromKey, hash, fmt.Errorf("unknown type: %d", ptype)
	}
	s := rlp.NewStream(bytes.NewReader(sigdata[1:]), 0)
	err = s.Decode(req)
	fmt.Println("Decode", req, fromKey, hash, err)
	return req, fromKey, hash, err
}

// Encode encodes a discovery packet.
// Encode는 discovery 패킷을 인코딩합니다.
func Encode(priv *ecdsa.PrivateKey, req Packet) (packet, hash []byte, err error) {
	b := new(bytes.Buffer)
	b.Write(headSpace)
	b.WriteByte(req.Kind())
	if err := rlp.Encode(b, req); err != nil {
		return nil, nil, err
	}
	packet = b.Bytes()
	sig, err := crypto.Sign(crypto.Keccak256(packet[headSize:]), priv)
	if err != nil {
		return nil, nil, err
	}
	copy(packet[macSize:], sig)
	// Add the hash to the front. Note: this doesn't protect the packet in any way.
	// 해시를 앞에 추가합니다. 참고: 이렇게 해도 패킷은 보호되지 않습니다.
	hash = crypto.Keccak256(packet[macSize:])
	copy(packet, hash)
	fmt.Println("Encode", packet, hash)
	return packet, hash, nil
}

// recoverNodeKey computes the public key used to sign the given hash from the signature.
// recoverNodeKey는 서명에서 지정된 해시를 서명하는 데 사용되는 공용 키를 계산합니다.
func recoverNodeKey(hash, sig []byte) (key Pubkey, err error) {
	pubkey, err := crypto.Ecrecover(hash, sig)
	if err != nil {
		return key, err
	}
	copy(key[:], pubkey[1:])
	fmt.Println("recoverNodeKey", key)
	return key, nil
}

// EncodePubkey encodes a secp256k1 public key.
// EncodePubkey는 secp256k1 공개 키를 인코딩합니다.
func EncodePubkey(key *ecdsa.PublicKey) Pubkey {
	var e Pubkey
	math.ReadBits(key.X, e[:len(e)/2])
	math.ReadBits(key.Y, e[len(e)/2:])
	fmt.Println("EncodePubkey", key)
	return e
}

// DecodePubkey reads an encoded secp256k1 public key.
// DecodePubkey는 인코딩된 secp256k1 공개 키를 읽습니다.
func DecodePubkey(curve elliptic.Curve, e Pubkey) (*ecdsa.PublicKey, error) {
	p := &ecdsa.PublicKey{Curve: curve, X: new(big.Int), Y: new(big.Int)}
	half := len(e) / 2
	p.X.SetBytes(e[:half])
	p.Y.SetBytes(e[half:])
	if !p.Curve.IsOnCurve(p.X, p.Y) {
		return nil, ErrBadPoint
	}
	fmt.Println("DecodePubkey", p)
	return p, nil
}
