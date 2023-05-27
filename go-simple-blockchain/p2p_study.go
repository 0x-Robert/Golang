package main

import (
	"bufio"
	"context"
	"crypto/rand"
	"fmt"
	"io"
	"log"
	mrand "math/rand"

	libp2p "github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/peerstore"
	crypto "github.com/libp2p/go-libp2p/core/crypto"
	host "github.com/libp2p/go-libp2p/core/host"
	ma "github.com/multiformats/go-multiaddr"
)

// random peer ID를 가진 LibP2P 호스트를 만듭니다.
func makeBasicHost(listenPort int, secio bool, randseed int64) (host.Host, error) {
	// randseed가 0이면 완벽한 랜덤값이 아닙니다. 예측가능한 값이 사용되어 같은 priv가 생성될 것입니다.
	var r io.Reader
	if randseed == 0 {
		r = rand.Reader
	} else {
		r = mrand.New(mrand.NewSource(randseed))
	}

	// 이 호스트의 key pair를 만듭니다.
	priv, _, err := crypto.GenerateKeyPairWithReader(crypto.RSA, 2048, r)
	if err != nil {
		return nil, err
	}

	// 옵션들.
	opts := []libp2p.Option{
		libp2p.ListenAddrStrings(fmt.Sprintf("/ip4/0.0.0.0/tcp/%d", listenPort)),
		libp2p.Identity(priv),
		libp2p.DisableRelay(),
	}

	// 호스트를 만들어 리턴합니다.
	return libp2p.New(context.Background(), opts...)
}

// {data}(cmd + payload)를 보냄
// p2p 에서는 peer ID를 이용하여 통신합니다.
func SendData(destPeerID string, data []byte) {
	peerID, err := peer.Decode(destPeerID)
	if err != nil {
		log.Panic(err)
	}

	// {ha} => {peerID} 의 Stream을 만듭니다.
	// 이 Stream은 {peerID}호스트의 steamHandler에 의해 처리될 것입니다.
	s, err := ha.NewStream(context.Background(), peerID, "/p2p/1.0.0")
	if err != nil {
		log.Printf("%s is not reachable\n", destPeerID)
		// TODO: 통신이 되지 않는 {peer}를 KnownNodes에서 삭제합니다.
		var updatedPeers []string

		// 통신이 되지 않는 {addr}를 KnownNodes에서 삭제합니다.
		for _, node := range KnownNodes {
			if node != destPeerID {
				updatedPeers = append(updatedPeers, node)
			}
		}

		KnownNodes = updatedPeers

		return
	}
	defer s.Close()

	_, err = s.Write(data)
	if err != nil {
		log.Println(err)
		return
	}
}

// {targetPeer}에게 {data}를 보냅니다.
// 1회성 host를 만들어 전송합니다.
func SendDataOnce(targetPeer string, data []byte) {
	host, err := libp2p.New(context.Background())
	if err != nil {
		log.Panic(err)
	}
	defer host.Close()
	ha = host

	destPeerID := addAddrToPeerstore(host, targetPeer)
	SendData(peer.Encode(destPeerID), data)
}

// 호스트의 0번째 주소를 알아옵니다.
func getHostAddress(_ha host.Host) string {
	// Build host multiaddress
	hostAddr, _ := ma.NewMultiaddr(fmt.Sprintf("/ipfs/%s", _ha.ID().Pretty()))

	// Now we can build a full multiaddress to reach this host
	// by encapsulating both addresses:
	addr := _ha.Addrs()[0]
	return addr.Encapsulate(hostAddr).String()
}

// Stream을 받았을 때 처리하는 핸들러 함수
func handleStream(s network.Stream) {
	// 일이 다 끝나면 stream을 종료합니다.
	defer s.Close()

	// Non blocking read/write를 위해 버퍼 스트림을 만듭니다.
	rw := bufio.NewReadWriter(bufio.NewReader(s), bufio.NewWriter(s))

	// connection 처리는 asynchronous하게 go routine으로 처리
	go HandleP2PConnection(rw)
}

// Host를 시작합니다.
func StartHost(listenPort int, minter string, secio bool, randseed int64, targetPeer string) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// minter의 주소를 global 변수에 저장.
	minterAddress = minter

	// {listenPort}가 nodeId로 쓰이게됩니다.
	nodeId := fmt.Sprintf("%d", listenPort)
	chain = blockchain.ContinueBlockChain(nodeId)
	go CloseDB(chain) // 하드웨어 인터럽트를 대기하고 있다가 안전하게 DB를 닫는 함수
	defer chain.Database.Close()

	// p2p host를 만듭니다.
	host, err := makeBasicHost(listenPort, secio, randseed)
	if err != nil {
		log.Panic(err)
	}
	// {host}를 전역변수 {ha}에 저장합니다.
	ha = host
	// {nodePeerId}: 이 노드의 peer ID 입니다.
	// 통신에 Peer Id 가 사용됩니다.
	nodePeerId = peer.Encode(host.ID())

	if len(KnownNodes) == 0 {
		// KnownNodes[0]는 자기 자신입니다.
		KnownNodes = append(KnownNodes, nodePeerId)
	}

	if targetPeer == "" {
		// listen 합니다.
		runListener(ctx, ha, listenPort, secio)
	} else {
		// listen하면서 listening하고 있는 서버에 접속합니다.
		runSender(ctx, ha, targetPeer)
	}

	// Wait forever
	select {}
}

// listening server를 구동합니다. (중앙 서버)
func runListener(ctx context.Context, ha host.Host, listenPort int, secio bool) {
	fullAddr := getHostAddress(ha)
	log.Printf("I am %s\n", fullAddr)

	// StreamHandler를 Set합니다.
	// {handleStream}은 stream을 받았을 때 불리는 핸들러 함수 입니다.
	// /p2p/1.0.0은 user-defined protocal 입니다.
	ha.SetStreamHandler("/p2p/1.0.0", handleStream)

	log.Printf("Now run \"go run main.go startp2p -dest %s\" on a different terminal\n", fullAddr)
}

// StreamHandler를 설정하고, {targetPeer}에게 Version 정보를 보냅니다.
func runSender(ctx context.Context, ha host.Host, targetPeer string) {
	fullAddr := getHostAddress(ha)
	log.Printf("I am %s\n", fullAddr)

	ha.SetStreamHandler("/p2p/1.0.0", handleStream)

	// targetPeer를 ha의 Peerstore에 저장하고 destination의 peerId를 받아옵니다.
	destPeerID := addAddrToPeerstore(ha, targetPeer)

	// {destPeerID}에게 {chain}의 Version을 보냅니다.
	SendVersion(peer.Encode(destPeerID), chain)
}

// peer의 {addr}를 받아 multiaddress로 파싱한 후 host의 peerstore에 저장합니다.
// 해당 정보로 peer ID를 알면 어떻게 통신해야하는 지 알 수 있습니다.
// peer의 ID를 반환합니다.
func addAddrToPeerstore(ha host.Host, addr string) peer.ID {
	// multiaddress로 파싱 후
	ipfsaddr, err := ma.NewMultiaddr(addr)
	if err != nil {
		log.Fatalln(err)
	}

	// multiaddress에서 Address와 PeerID 정보를 알아옵니다.
	info, err := peer.AddrInfoFromP2pAddr(ipfsaddr)
	if err != nil {
		log.Fatalln(err)
	}

	// LibP2P가 참고할 수 있도록
	// Peer ID와 address를 peerstore에 저장합니다.
	ha.Peerstore().AddAddrs(info.ID, info.Addrs, peerstore.PermanentAddrTTL)
	return info.ID
}