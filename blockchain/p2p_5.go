package main

import (
	"bufio"
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	mrand "math/rand"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/davecgh/go-spew/spew"
	libp2p "github.com/libp2p/go-libp2p"
	crypto "github.com/libp2p/go-libp2p/core/crypto"
	host "github.com/libp2p/go-libp2p/core/host"
	net "github.com/libp2p/go-libp2p/core/network"
	peer "github.com/libp2p/go-libp2p/core/peer"
	pstore "github.com/libp2p/go-libp2p/core/peerstore"
	ma "github.com/multiformats/go-multiaddr"
	// golog "github.com/ipfs/go-log/v2"
)

type Block struct {
    Index     int   
    Timestamp string
    BPM       int   
    Hash      string
    PrevHash  string
}
var Blockchain []Block
var mutex = &sync.Mutex{}



func isBlockValid(newBlock, oldBlock Block) bool {
    if oldBlock.Index+1 != newBlock.Index {
        return false
    }

    if oldBlock.Hash != newBlock.PrevHash {
        return false
    }
    
    if calculateHash(newBlock) != newBlock.Hash {
        return false
    }
    
    return true
}

// SHA256 hashing
func calculateHash(block Block) string {
    record := strconv.Itoa(block.Index) + block.Timestamp + strconv.Itoa(block.BPM) + block.PrevHash
    h := sha256.New()
    h.Write([]byte(record))
    hashed := h.Sum(nil)
    return hex.EncodeToString(hashed)
}

// create a new block using previous block's hash
func generateBlock(oldBlock Block, BPM int) Block {

    var newBlock Block
    
    t := time.Now()
    
    newBlock.Index = oldBlock.Index + 1
    newBlock.Timestamp = t.String()
    newBlock.BPM = BPM
    newBlock.PrevHash = oldBlock.Hash
    newBlock.Hash = calculateHash(newBlock)
    
    return newBlock
}

//listenPort는 다른 perr가 접근할 수 있도록 하는 포트
//secio는 secure input/output의 약자 안전하게 데이터 스트림 사용할지 체크
//randSeed는 호스트의 임의주소를 생성할지를 결정하는 부가적인 인자 
func makeBasicHost(listenPort int, secio bool, randseed int64) (host.Host, error) {

    // If the seed is zero, use real cryptographic randomness. Otherwise, use a
    // deterministic randomness source to make generated keys stay the same
    // across multiple runs
	//randseed가 0이면 완벽한 랜덤값이 아님 
    var r io.Reader
    if randseed == 0 {
        r = rand.Reader 
    } else {
		//randseed 사용하는지 확인하고 호스트 키 생성 
        r = mrand.New(mrand.NewSource(randseed))
    }
    
    // Generate a key pair for this host. We will use it
    // to obtain a valid host ID.
    priv, _, err := crypto.GenerateKeyPairWithReader(crypto.RSA, 2048, r)
    if err != nil {
        return nil, err
    }
    
    opts := []libp2p.Option{
        libp2p.ListenAddrStrings(fmt.Sprintf("/ip4/127.0.0.1/tcp/%d", listenPort)),
        libp2p.Identity(priv),
		libp2p.DisableRelay(),
    }

    
	if !secio {
		opts = append(opts, libp2p.NoSecurity)
	}

	return libp2p.New(opts...)
    // //basicHost, err := libp2p.New(context.Background(), opts)
	basicHost, err := libp2p.New(context.Background())

    if err != nil {
        return nil, err
    }
    
    //Build host multiaddress
    hostAddr, _ := ma.NewMultiaddr(fmt.Sprintf("/ipfs/%s", basicHost.ID().Pretty()))
    
    // // Now we can build a full multiaddress to reach this host
    // by encapsulating both addresses:
    addr := basicHost.Addrs()[0]
    fullAddr := addr.Encapsulate(hostAddr)
    log.Printf("I am %s\n", fullAddr)
    if secio {
        log.Printf("Now run \"go run main.go -l %d -d %s -secio\" on a different terminal\n", listenPort+1, fullAddr)
    } else {
        log.Printf("Now run \"go run main.go -l %d -d %s\" on a different terminal\n", listenPort+1, fullAddr)
    }
    
    return basicHost, nil
}

//다른 노드에서 현재 호스트에 연결 후 블록 생성 후 블록체인에 추가할지 결정하는 코드 
func handleStream(s net.Stream) {
    

	log.Println("Got a new stream!")


	// Create a buffer stream for non blocking read and write.
	rw := bufio.NewReadWriter(bufio.NewReader(s), bufio.NewWriter(s))

	//고루틴으로 다음 함수들 실행 
	go readData(rw)
	go writeData(rw)


	// stream 's' will stay open until you close it (or the other side closes it).
}


func readData(rw *bufio.ReadWriter) {
    

	for {
		str, err := rw.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}


		if str == "" {
			return
		}
		if str != "\n" {


			chain := make([]Block, 0)
			if err := json.Unmarshal([]byte(str), &chain); err != nil {
				log.Fatal(err)
			}


			mutex.Lock()//동시실행하는 것을 막아주는 mutex 락
			//채인이 블록체인보다 길면 블록체인 변수에 채인값을 넣는다.
			//이후 바이트와 err 변수에 블록체인에 MarshalIndent함수를 적용한 값을 저장한다. 
			if len(chain) > len(Blockchain) {
				Blockchain = chain
				bytes, err := json.MarshalIndent(Blockchain, "", "  ")
				if err != nil {


					log.Fatal(err)
				}
				// Green console color:     \x1b[32m
				// Reset console color:     \x1b[0m
				fmt.Printf("\x1b[32m%s\x1b[0m> ", string(bytes))
			}
			mutex.Unlock()//동시실행하는 것을 막아주는 mutex 언락
		}
	}
}

func writeData(rw *bufio.ReadWriter) {
    

	go func() {
		for {
			//5초 슬립
			time.Sleep(5 * time.Second)
			//뮤텍스 락, 동시실행 막아줌
			mutex.Lock()
			//블록체인 데이터를 json화한 값을 bytes에 저장 
			bytes, err := json.Marshal(Blockchain)
			if err != nil {
				log.Println(err)
			}
			//뮤텍스 언락, 동시실행 막아줌
			mutex.Unlock()

			//뮤텍스 락, 동시실행 막아줌
			mutex.Lock()
			rw.WriteString(fmt.Sprintf("%s\n", string(bytes)))
			rw.Flush()
			//뮤텍스 언락, 동시실행 막아줌
			mutex.Unlock()


		}
	}()


	stdReader := bufio.NewReader(os.Stdin)


	for {
		fmt.Print("> ")
		sendData, err := stdReader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}


		sendData = strings.Replace(sendData, "\n", "", -1)
		bpm, err := strconv.Atoi(sendData)
		if err != nil {
			log.Fatal(err)
		}
		newBlock := generateBlock(Blockchain[len(Blockchain)-1], bpm)


		if isBlockValid(newBlock, Blockchain[len(Blockchain)-1]) {
			mutex.Lock()
			Blockchain = append(Blockchain, newBlock)
			mutex.Unlock()
		}


		bytes, err := json.Marshal(Blockchain)
		if err != nil {
			log.Println(err)
		}


		spew.Dump(Blockchain)


		mutex.Lock()
		rw.WriteString(fmt.Sprintf("%s\n", string(bytes)))
		rw.Flush()
		mutex.Unlock()
	}


}


func main() {
	t := time.Now()
	genesisBlock := Block{}
	genesisBlock = Block{0, t.String(), 0, calculateHash(genesisBlock), ""}


	Blockchain = append(Blockchain, genesisBlock)


	// LibP2P code uses golog to log messages. They log with different
	// string IDs (i.e. "swarm"). We can control the verbosity level for
	// all loggers with:
	//golog.SetAllLoggers(gologging.INFO) // Change to DEBUG for extra info


	// Parse options from the command line
	listenF := flag.Int("l", 0, "wait for incoming connections")
	target := flag.String("d", "", "target peer to dial")
	secio := flag.Bool("secio", false, "enable secio")
	seed := flag.Int64("seed", 0, "set random seed for id generation")
	flag.Parse()


	if *listenF == 0 {
		log.Fatal("Please provide a port to bind on with -l")
	}


	// Make a host that listens on the given multiaddress
	ha, err := makeBasicHost(*listenF, *secio, *seed)
	if err != nil {
		log.Fatal(err)
	}


	if *target == "" {
		log.Println("listening for connections")
		// Set a stream handler on host A. /p2p/1.0.0 is
		// a user-defined protocol name.
		ha.SetStreamHandler("/p2p/1.0.0", handleStream)


		select {} // hang forever
		/**** This is where the listener code ends ****/
	} else {
		ha.SetStreamHandler("/p2p/1.0.0", handleStream)


		// The following code extracts target's peer ID from the
		// given multiaddress
		ipfsaddr, err := ma.NewMultiaddr(*target)
		if err != nil {
			log.Fatalln(err)
		}


		pid, err := ipfsaddr.ValueForProtocol(ma.P_IPFS)
		if err != nil {
			log.Fatalln(err)
		}


		// peerid, err := peer.IDB58Decode(pid)
		// if err != nil {
		// 	log.Fatalln(err)
		// }

		peerid, err := peer.Decode(pid)
		if err != nil {
			log.Fatalln(err)
		}

		// Decapsulate the /ipfs/<peerID> part from the target
		// /ip4/<a.b.c.d>/ipfs/<peer> becomes /ip4/<a.b.c.d>
		targetPeerAddr, _ := ma.NewMultiaddr(
			// fmt.Sprintf("/ipfs/%s", peer.IDB58Encode(peerid)))
			fmt.Sprintf("/ipfs/%s", peer.Encode(peerid)))
		targetAddr := ipfsaddr.Decapsulate(targetPeerAddr)


		// We have a peer ID and a targetAddr so we add it to the peerstore
		// so LibP2P knows how to contact it
		ha.Peerstore().AddAddr(peerid, targetAddr, pstore.PermanentAddrTTL)


		log.Println("opening stream")
		// make a new stream from host B to host A
		// it should be handled on host A by the handler we set above because
		// we use the same /p2p/1.0.0 protocol
		s, err := ha.NewStream(context.Background(), peerid, "/p2p/1.0.0")
		if err != nil {
			log.Fatalln(err)
		}
		// Create a buffered stream so that read and writes are non blocking.
		rw := bufio.NewReadWriter(bufio.NewReader(s), bufio.NewWriter(s))


		// Create a thread to read and write data.
		go writeData(rw)
		go readData(rw)


		select {} // hang forever


	}
}

