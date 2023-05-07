package main

import (
	"bufio"         //버퍼 인풋 아웃풋 패키지
	"context"       //컨텍스트 관련 패키지, 고루틴을 관리하거나 네트워크의 취소등의 작업을 할 때 필요함
	"crypto/rand"   //암호학 난수 생성 패키지?
	"crypto/sha256" //sha256 패키지
	"encoding/hex"  //hex encoding
	"encoding/json" //json encoding 패키지
	"flag"          //flag
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
)

//블록 구조체 정의

type Block struct {
	Index     int //블록 인덱스 
	Timestamp string //타임스탬프, 시간 
	BPM       int // BPM 분당 비트 수 ? 맥박? 
	Hash      string // 해쉬 
	PrevHash  string // 이전 해쉬 
}

// Blockchain is a series of validated Blocks
// 블록체인 변수는 배열 값 
var Blockchain []Block

//동시실행을 방지하는 mutex
var mutex = &sync.Mutex{}

// make sure block is valid by checking index, and comparing the hash of the previous block
// 블록 유효성 검사하는 함수
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
// 블록의 인덱스, 타임스탬프, BPM, 이전 해시값을  sha 256 해싱값 계산해서 리턴해주는 함수 
func calculateHash(block Block) string {
	record := strconv.Itoa(block.Index) + block.Timestamp + strconv.Itoa(block.BPM) + block.PrevHash
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

// create a new block using previous block's hash
// 블록을 생성하는 함수 
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

// makeBasicHost creates a LibP2P host with a random peer ID listening on the
// given multiaddress. It will use secio if secio is true.
// makeBasicHost는 수신 대기하는 임의의 피어 ID로 LibP2P 호스트를 생성합니다 
// 주소가 여러 개 지정되었습니다. secio가 참이면 secio를 사용합니다
func makeBasicHost(listenPort int, secio bool, randseed int64) (host.Host, error) {

	// If the seed is zero, use real cryptographic randomness. Otherwise, use a
	// deterministic randomness source to make generated keys stay the same
	// across multiple runs
    // 시드가 0이면 실제 암호화 랜덤성을 사용합니다. 그렇지 않으면 결정론적 랜덤성 소스를 사용하여 생성된 키가 여러 실행에서 동일하게 유지되도록 합니다
	//rand.Seed(seed int64) 함수 : 랜덤 시드 설정
	var r io.Reader
	if randseed == 0 {
		r = rand.Reader
	} else {
		//r에는 랜덤 값을 설정 
		//NewSource는 지정된 값으로 시드된 새 의사 임의 소스를 반환합니다. 
		//최상위 함수에서 사용하는 기본 소스와 달리 이 소스는 여러 고루틴에서 동시에 사용하기에 안전하지 않습니다. 반환된 소스는 소스 64를 구현합니다.
		r = mrand.New(mrand.NewSource(randseed))
	}

	// Generate a key pair for this host. We will use it
	// to obtain a valid host ID.
	//GenerateKeyPairWithReader는 지정된 유형 및 비트 크기의 키 쌍을 반환합니다
	/*
	func GenerateKeyPairWithReader(typ, bits int, src io.Reader) (PrivKey, PubKey, error) {
	switch typ {
	case RSA:
		return GenerateRSAKeyPair(bits, src)
	case Ed25519:
		return GenerateEd25519Key(src)
	case Secp256k1:
		return GenerateSecp256k1Key(src)
	case ECDSA:
		return GenerateECDSAKeyPair(src)
	default:
		return nil, nil, ErrBadKeyType
	}
}
	*/
	//타입, 비트, 인풋아웃풋 리더 
	priv, _, err := crypto.GenerateKeyPairWithReader(crypto.RSA, 2048, r)
	if err != nil {
		return nil, err
	}

	/*
	
	func ListenAddrStrings(s ...string) Option {
	return func(cfg *Config) error {
		for _, addrstr := range s {
			a, err := ma.NewMultiaddr(addrstr)
			if err != nil {
				return err
			}
			cfg.ListenAddrs = append(cfg.ListenAddrs, a)
		}
		return nil
	}
}


    Identity configures libp2p to use the given private key to identify itself.
	func Identity(sk crypto.PrivKey) Option {
		return func(cfg *Config) error {
			if cfg.PeerKey != nil {
				return fmt.Errorf("cannot specify multiple identities")
			}

			cfg.PeerKey = sk
			return nil
		}
	}

	*/
	//libp2p.Option ==> 옵션은 libp2p 생성자에 제공할 수 있는 libp2p 구성 옵션입니다.
	opts := []libp2p.Option{
		//ListenAddrStrings는 지정된(파싱되지 않음) 수신을 위해 libp2p를 구성합니다
		libp2p.ListenAddrStrings(fmt.Sprintf("/ip4/127.0.0.1/tcp/%d", listenPort)),
		// ID는 지정된 개인 키를 사용하여 자신을 식별하도록 libp2p를 구성합니다.
		libp2p.Identity(priv),
	}

	//func append(slice []Type, elems ...Type) []Type

	if !secio {
		//보안 없음은 모든 전송 보안을 완전히 비활성화하는 옵션입니다.

		opts = append(opts, libp2p.NoSecurity)
	}

	//반환된 libp2p 노드를 중지/종료하려면 사용자가 전달된 컨텍스트를 취소하고 반환된 호스트에서 '닫기'를 호출해야 합니다.
	/*
	func New(opts ...Option) (host.Host, error) {
	return NewWithoutDefaults(append(opts, FallbackDefaults)...)
}
	*/
	basicHost, err := libp2p.New(opts...)
	if err != nil {
		return nil, err
	}

	// Build host multiaddress
	/*
	func NewMultiaddr(s string) (a Multiaddr, err error) {
	defer func() {
		if e := recover(); e != nil {
			log.Printf("Panic in NewMultiaddr on input %q: %s", s, e)
			err = fmt.Errorf("%v", e)
		}
	}()
	b, err := stringToBytes(s)
	if err != nil {
		return nil, err
	}
	return &multiaddr{bytes: b}, nil
}
	func (id ID) Pretty() string {
		return id.String()
	}
	
	*/
	hostAddr, _ := ma.NewMultiaddr(fmt.Sprintf("/ipfs/%s", basicHost.ID().Pretty()))

	// Now we can build a full multiaddress to reach this host
	// by encapsulating both addresses:
	// Returns the listen addresses of the Host
	// Addrs() []ma.Multiaddr
	addr := basicHost.Addrs()[0]

	// Encapsulate wraps this Multiaddr around another. For example:
	//
	//      /ip4/1.2.3.4 encapsulate /tcp/80 = /ip4/1.2.3.4/tcp/80
	//
	//	Encapsulate(Multiaddr) Multiaddr
	fullAddr := addr.Encapsulate(hostAddr)
	log.Printf("I am  fullAddr %s\n", fullAddr)
	if secio {
		log.Printf("Now run \"go run p2p_5.go -l %d -d %s -secio\" on a different terminal\n", listenPort+1, fullAddr)
	} else {
		log.Printf("Now run \"go run p2p_5.go -l %d -d %s\" on a different terminal\n", listenPort+1, fullAddr)
	}

	return basicHost, nil
}


//스트림에 대한 것을 다루는 함수고
//스트림이란  두 에이전트 사이의 양방향 채널을 나타냅니다
func handleStream(s net.Stream) {

	log.Println("Got a new stream!")

	// Create a buffer stream for non blocking read and write.
	// 읽기 및 쓰기를 차단하지 않는 버퍼 스트림을 만듭니다.
	/*
	type ReadWriter struct {
	*Reader
	*Writer
}

	NewReadWriter allocates a new ReadWriter that dispatches to r and w.
	func NewReadWriter(r *Reader, w *Writer) *ReadWriter {
		return &ReadWriter{r, w}
	}
	*/
	rw := bufio.NewReadWriter(bufio.NewReader(s), bufio.NewWriter(s))

	//새로운 ReadWrite를 생성하고 Go루틴을 이용하여 read, write를 분리하여 처리한다. 
	go readData(rw)
	go writeData(rw)

	// stream 's' will stay open until you close it (or the other side closes it).
	//스트림 's'는 사용자가 닫을 때까지(또는 다른 쪽이 닫을 때까지) 열려 있습니다.
}

func readData(rw *bufio.ReadWriter) {
//블록체인을 언제든지 전송받아 읽을 수 있도록 무한루프 안에 로직을 작성하기 
// 전송받은 블록체인 문자열을 ReadString으로 파싱 
// 만약에 문자열이 비어있지 않다면 Unmarshal 디코딩하기 
	for {
		//rw를 읽기 / 줄 바꾸기를 통해 읽기 
		str, err := rw.ReadString('\n')
		//에러가 있으면 log.Fatal로 출력 
		if err != nil {
			log.Fatal(err)
		}

		//str이 공백이면  리턴 
		if str == "" {
			return
		}


		if str != "\n" {
			//채널 생성 
			chain := make([]Block, 0)
			//json Unmarshal 함수 호출 (json 데이터를 원본 데이터로 변환 )
			if err := json.Unmarshal([]byte(str), &chain); err != nil {
				log.Fatal(err)
			}

			mutex.Lock()
			//동시실행 방지 mutex lock
			//체인이 더 길경우 블록체인은 긴 체인 값이 된다.  
			if len(chain) > len(Blockchain) {
				Blockchain = chain
				//바이트 값을 MarshalIndent 함수를 호출한 값에 담는다 Json데이터로 만들고 들여쓰기도 적용함 
				bytes, err := json.MarshalIndent(Blockchain, "", "  ")
				if err != nil {

					log.Fatal(err)
				}
				// Green console color: 	\x1b[32m
				// Reset console color: 	\x1b[0m
				fmt.Printf("\x1b[32m%s\x1b[0m> ", string(bytes))
			}
			//동시실행방지 mutex 언락 
			mutex.Unlock()
		}
	}
}

//JSON 형식으로 Marshal 인코딩 한 후 fmt.printf 새로운 블록체인을 보기 쉽게 출력합니다.
func writeData(rw *bufio.ReadWriter) {

	//매 5초마다 각 peer들에게 업데이트된 블록체인을 알려주는 고루틴 실행 
	go func() {
		for {
			time.Sleep(5 * time.Second)
			//동시실행 방지 mutex lock
			mutex.Lock()
			//Json 형식으로 Marshal 인코딩 후 새로운 블록체인을 프린트 
			bytes, err := json.Marshal(Blockchain)
			if err != nil {
				log.Println(err)
			}
			//동시실행 방지 mutex unlock
			mutex.Unlock()

			//동시실행 방지 mutex lock 
			mutex.Lock()
			// 쓰기 인스턴스로 버퍼에 string(bytes) 쓰기
			//rw.WriteString을 사용하여 연결된 peer에 업데이트 된 블록체인을 전송 합니다.
			rw.WriteString(fmt.Sprintf("%s\n", string(bytes)))
			//버퍼의 내용을 파일에 저장
			rw.Flush()
			//동시실행 방지 mutex unlock 
			mutex.Unlock()

		}
	}()
		
	//stdReader 변수에 버퍼 NewReader 저장 
	//Bufio.NewReader를 가지고 새로운 reader를 생성하면 stdin(콘솔 입력)를 통해 BPM을 입력 할 수 있습니다. 
	stdReader := bufio.NewReader(os.Stdin)

	//반복순회 
	for {

		fmt.Print("> ")
		//sendData 변수 저장 
		sendData, err := stdReader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		
		sendData = strings.Replace(sendData, "\n", "", -1)
		//bpm에 strconv.Atoi(sendData)값저장
		//strconv.Atoi >> 숫자로 이러우진 문자열을 숫자로 변환 
		bpm, err := strconv.Atoi(sendData)
		if err != nil {
			log.Fatal(err)
		}
		
		newBlock := generateBlock(Blockchain[len(Blockchain)-1], bpm)

		//블록 유효성 검사 
		if isBlockValid(newBlock, Blockchain[len(Blockchain)-1]) {
			//동시실행방지 mutex lock 
			mutex.Lock()
			//블록체인에 뉴블록 저장 
			Blockchain = append(Blockchain, newBlock)
			//동시실행방지 mutex unlock
			mutex.Unlock()
		}

		//바이트에 json.Marshal함수를 통해 블록체인 데이터를 json화하고 저장 
		bytes, err := json.Marshal(Blockchain)
		if err != nil {
			log.Println(err)
		}
		//디버그를 위한 spew 사용 
		spew.Dump(Blockchain)

		//동시실행방지 mutex lock 
		mutex.Lock()

		//rw 버퍼에 다음 항목들 쓰기 실행 
		//rw.WriteString을 사용하여 연결된 peer에 업데이트 된 블록체인을 전송 합니다.
		rw.WriteString(fmt.Sprintf("%s\n", string(bytes)))

		//rw 버퍼에 저장 
		rw.Flush()

		//동시실행방지 mutex unlock 
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

	// LibP2P 코드는 golog를 사용하여 메시지를 기록합니다. 서로 다른 로그를 기록합니다	
	// 문자열 ID(예: "swarm"). 우리는 다음에 대한 장황한 수준을 제어할 수 있다	
	// 모든 로거:	//golog.모든 로거를 설정합니다(로그인).정보) 
	// 추가 정보를 보려면 DEBUG로 변경	
	// 명령줄에서 옵션 구문 분석
	

	//flag 패키지는 command-line flag, 즉, 인자들을 파싱할 수 있게 도와주는 패키지다. 
	//일단 flag 패키지에 있는 func Int64 함수를 살펴보자. 이 함수는 int64 포인터를 반환하는 함수다.
	//maxValue := flag.Int64("max", 10, "Defines maximum value")
	// Secio 보안 스트림을 사용여부를 선택하는 flag값입니다.
	// Target은 다른 호스트의 주소를 나타냅니다.
	// Listen는 이 포트를 통해 다른 peer가 접속 하도록하는데 사용됩니다.
	// Seed는 임의적인 주소를 만들지 선택하는 값입니다.
	listenF := flag.Int("l", 0, "wait for incoming connections")
	target := flag.String("d", "", "target peer to dial")
	secio := flag.Bool("secio", false, "enable secio")
	seed := flag.Int64("seed", 0, "set random seed for id generation")
	yong := flag.String("y", "", "yongari test ")
	fmt.Println("yong",yong)
	flag.Parse()

	if *listenF == 0 {
		log.Fatal("Please provide a port to bind on with -l")
	}

	// Make a host that listens on the given multiaddress
	// 주어진 멀티 address에서 수신 대기하는 호스트로 만드는 함수 호출 
	ha, err := makeBasicHost(*listenF, *secio, *seed)

	//에러가 nil값이 아니면 로그로 출력 
	if err != nil {
		log.Fatal(err)
	}

	if *target == "" {
		log.Println("listening for connections")
		// Set a stream handler on host A. /p2p/1.0.0 is a user-defined protocol name.
		// 호스트 A에 스트림 핸들러를 설정합니다. /p2p/1.0.0은 사용자 정의 프로토콜 이름.
		ha.SetStreamHandler("/p2p/1.0.0", handleStream)

		select {} // hang forever
		//영원히 기달리다? 
		/**** This is where the listener code ends ****/
	} else {
		ha.SetStreamHandler("/p2p/1.0.0", handleStream)

		// The following code extracts target's peer ID from the
		// given multiaddress
		// 다음 코드는 대상의 피어 ID를 추출합니다
		// 지정된 복수 주소

		/*
		NewMultiaddr parses and validates an input string, returning a *Multiaddr
		func NewMultiaddr(s string) (a Multiaddr, err error) {
			defer func() {
				if e := recover(); e != nil {
					log.Printf("Panic in NewMultiaddr on input %q: %s", s, e)
					err = fmt.Errorf("%v", e)
				}
			}()
			b, err := stringToBytes(s)
			if err != nil {
				return nil, err
			}
			return &multiaddr{bytes: b}, nil
		}
		*/
		ipfsaddr, err := ma.NewMultiaddr(*target)
		if err != nil {
			log.Fatalln(err)
		}
		//	ValueForProtocol(code int) (string, error)
		pid, err := ipfsaddr.ValueForProtocol(ma.P_IPFS)
		if err != nil {
			log.Fatalln(err)
		}

		//peerid, err := peer.IDB58Decode(pid)
		peerid, err := peer.Decode(pid)
		if err != nil {
			log.Fatalln(err)
		}

		// Decapsulate the /ipfs/<peerID> part from the target
		// /ip4/<a.b.c.d>/ipfs/<peer> becomes /ip4/<a.b.c.d>
		// /ipfs/<peer 캡슐화 해제ID > 대상에서 분리
		// /ip4/<a.b.c.d>/ipfs/<peer>가 /ip4/<a.b.c.d>가 됩니다
		targetPeerAddr, _ := ma.NewMultiaddr(
			//fmt.Sprintf("/ipfs/%s", peer.IDB58Encode(peerid)))
			fmt.Sprintf("/ipfs/%s", peer.Encode(peerid)))
		
			//Decapsulate(Multiaddr) Multiaddr
		targetAddr := ipfsaddr.Decapsulate(targetPeerAddr)

		// We have a peer ID and a targetAddr so we add it to the peerstore
		// so LibP2P knows how to contact it
		// 피어 ID와 targetAddr이 있으므로 피어 저장소에 추가합니다
		// 그래서 LibP2P는 그것에 연락하는 방법을 안다
		//Peerstore() peerstore.Peerstore
		ha.Peerstore().AddAddr(peerid, targetAddr, pstore.PermanentAddrTTL)

		log.Println("opening stream")
		// make a new stream from host B to host A
		// it should be handled on host A by the handler we set above because
		// we use the same /p2p/1.0.0 protocol

		// 호스트 B에서 호스트 A로 새로운 스트림을 만들다
		// 호스트 A에서 우리가 위에서 설정한 핸들러에 의해 처리되어야 한다
		// 동일한 /p2p/1.0.0 프로토콜을 사용합니다
		//NewStream(ctx context.Context, p peer.ID, pids ...protocol.ID) (network.Stream, error)
		s, err := ha.NewStream(context.Background(), peerid, "/p2p/1.0.0")
		if err != nil {
			log.Fatalln(err)
		}
		// Create a buffered stream so that read and writes are non blocking.
		// 읽기 및 쓰기가 차단되지 않도록 버퍼링된 스트림을 만듭니다.
		rw := bufio.NewReadWriter(bufio.NewReader(s), bufio.NewWriter(s))

		// Create a thread to read and write data.
		// 데이터를 읽고 쓸 스레드를 만듭니다.
		go writeData(rw)
		go readData(rw)

		select {} // hang forever

	}
}
//go run main.go -l 10000 -secio