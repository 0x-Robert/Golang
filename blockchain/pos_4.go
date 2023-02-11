package main

import (
	"bufio"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/joho/godotenv"
)

//블록에 필요한 정보를 담을 수 있는 구조체
type Block struct {
    Index     int
    Timestamp string
    BPM       int
    Hash      string
    PrevHash  string
    Validator string
}

// Blockchain is a series of validated Blocks
//블록체인 
var Blockchain []Block
//블록을 생성하는 노드를 선택하기 전에 블록을 보관하는 임시 저장소 
var tempBlocks []Block

// 블록들의 채널, 새로운 블록을 제출하는 각 노드들은 이 채널로 보냄 
// candidateBlocks handles incoming blocks for validation
var candidateBlocks = make(chan Block)

// TCP 서버에서 업데이트 된 블록체인을 보내주는 채널이다. 
// announcements broadcasts winning validator to all nodes
// make를 사용해서 초기화해줘야하는데 슬라이스와 맵과 채널의 경우 make를 이용한다. 
var announcements = make(chan string)

//데아터가 중복되는 것을 방지함 
// 상호 배제(mutual exclusion)라고도 하며 여러 스레드(고루틴)에서 공유되는 데이터를 보호할 때 주로 사용합니다.
var mutex = &sync.Mutex{}

//노드가 가지고 있는 토큰 수를 나타냄
// validators keeps track of open validators and balances
var validators = make(map[string]int)


	//문자열을 sha256 해시화하여 리턴 
	// calculateHash is a simple SHA256 hashing function
	func calculateHash(s string) string {
		h := sha256.New()
		fmt.Println("h",h)
		h.Write([]byte(s))
		hashed := h.Sum(nil)
		fmt.Println("hashed",hashed)
		return hex.EncodeToString(hashed)
    }
    
	//블록안의 데이터가 각 문자열을 연결하면서 해시화 함
    //calculateBlockHash returns the hash of all block information
    func calculateBlockHash(block Block) string {
		record := string(block.Index) + block.Timestamp + string(block.BPM) + block.PrevHash
		return calculateHash(record)
    }

	//새롭게 블록을 생성하는 함수 
	func generateBlock(oldBlock Block, BPM int, address string) (Block, error) {


		var newBlock Block
	
	
		t := time.Now()
	
	
		newBlock.Index = oldBlock.Index + 1
		newBlock.Timestamp = t.String()
		newBlock.BPM = BPM
		newBlock.PrevHash = oldBlock.Hash
		newBlock.Hash = calculateBlockHash(newBlock)
		newBlock.Validator = address
	
	
		return newBlock, nil
	}

	//블록의 유효성을 검증하는 함수 
	func isBlockValid(newBlock, oldBlock Block) bool {
		if oldBlock.Index +1 != newBlock.Index {
			return false
		}
	
		if oldBlock.Hash != newBlock.PrevHash {
			return false
		}
	
		//블록의 해시를 계산해서 유효성을 검증함 
		if calculateBlockHash(newBlock) != newBlock.Hash {
			return false
		}
	
	
		return true
	}

	//connection 핸들러
	func handleConn(conn net.Conn) {
        defer conn.Close()
    
		//goroutine을 통해 announcements를 msg로 넣는다
		// 채널 값을 msg에 넣는다. 
        go func() {
            for {
                msg := <-announcements
				
                io.WriteString(conn, msg)
            }
        }()
        // validator address
        var address string
    

        // allow user to allocate number of tokens to stake
        // the greater the number of tokens, the greater chance to forging a new block
		// 토큰의 개수를 입력받는다. 
        io.WriteString(conn, "Enter token balance:")
        scanBalance := bufio.NewScanner(conn)
		// NewScanner() 메서드는 os.Stdin으로부터 문자를 읽어오는 Scanner를 생성, 반환합니다. 
		// Scanner는 Scan() 메서드와 Text() 메서드를 이용해 문자를 읽고, 변수에 저장할 수 있습니다. 다음의 예제를 살펴보겠습니다.
        for scanBalance.Scan() {
			//func Atoi(s string) (i int, err error): 숫자로 이루어진 문자열을 숫자로 변환
            balance, err := strconv.Atoi(scanBalance.Text())

            if err != nil {
                log.Printf("%v not a number: %v", scanBalance.Text(), err)
                return
            }
            t := time.Now()
            address = calculateHash(t.String())
            validators[address] = balance
            fmt.Println("validators", validators)
            break
        }
    

        io.WriteString(conn, "\nEnter a new BPM:")
    

        scanBPM := bufio.NewScanner(conn)
    

        go func() {
            for {
                // take in BPM from stdin and add it to blockchain after conducting necessary validation
                for scanBPM.Scan() {
                    bpm, err := strconv.Atoi(scanBPM.Text())
                    // if malicious party tries to mutate the chain with a bad input, delete them as a validator and they lose their staked tokens
                    if err != nil {
                        log.Printf("%v not a number: %v", scanBPM.Text(), err)
                        delete(validators, address) 
                        conn.Close()
                    }
    

                    mutex.Lock()
                    oldLastIndex := Blockchain[len(Blockchain)-1]
                    mutex.Unlock()
    

                    // create newBlock for consideration to be forged
					//새로운 블록을 생성함 
                    newBlock, err := generateBlock(oldLastIndex, bpm, address)
                    if err != nil {
                        log.Println(err)
                        continue
                    }
                    if isBlockValid(newBlock, oldLastIndex) {
						//블록 유효성이 검증되면 candidateBlock 채널로 새로운 블록을 보내야함 
                        candidateBlocks <- newBlock
                    }
                    io.WriteString(conn, "\nEnter a new BPM:")
                }
            }
        }()
    

        // simulate receiving broadcast
        for {
            time.Sleep(time.Minute)
            mutex.Lock()
            output, err := json.Marshal(Blockchain)
            mutex.Unlock()
            if err != nil {
                log.Fatal(err)
            }
            io.WriteString(conn, string(output)+"\n")
        }

	}

		func pickWinner() {
			//30초마다 동작
			time.Sleep(30 * time.Second)
			mutex.Lock()
			//temp에 임시 블록들 값 넣기 
			temp := tempBlocks
			mutex.Unlock()
		
			//lotteryPool 배열 생성 
			//선택된 노드의 주소를 보관하는 lotteryPool을 만들기 
			lotteryPool := []string{}

			//temp의 길이가 양수일 때 
			//실제 제출하는 새로운 블록이 있는지 확인 
			if len(temp) > 0 {
		
		
				// slightly modified traditional proof of stake algorithm
				// from all validators who submitted a block, weight them by the number of staked tokens
				// in traditional proof of stake, validators can participate without submitting a block to be forged
			OUTER:
				for _, block := range temp {
					// if already in lottery pool, skip
					// 만약 lottery 풀에 노드가 있는지 체크 있으면 반복문 나가기 
					for _, node := range lotteryPool {
						if block.Validator == node {
							continue OUTER
						}
					}
		
		
					// lock list of validators to prevent data race
					mutex.Lock()
					//벨리데이터 변수설정해서 벨리데이터 값을 넣기 
					setValidators := validators
					mutex.Unlock()
		
					//temp안에 있는 블록데이터가 신뢰할 수 있는 노드에서 생성된 것인지 확인하고 validator 안에 주소가 있는지 확인 
					k, ok := setValidators[block.Validator]
					//만약 validators 안에 주소가 있으면 lotteryPool에 추가함  
					if ok {
						for i := 0; i < k; i++ {
							lotteryPool = append(lotteryPool, block.Validator)
						}
					}
				}
		
		
				// randomly pick winner from lottery pool

				s := rand.NewSource(time.Now().Unix())
				r := rand.New(s)
				//선택된 노드 
				lotteryWinner := lotteryPool[r.Intn(len(lotteryPool))]
		
		
				// add block of winner to blockchain and let all the other nodes know
				for _, block := range temp {
					if block.Validator == lotteryWinner {
						mutex.Lock()
						Blockchain = append(Blockchain, block)
						mutex.Unlock()
						for _ = range validators {
							//선택된 노드를 알려줌 
							announcements <- "\nwinning validator: " + lotteryWinner + "\n"
						}
						break
					}
				}
			}
		
		
			mutex.Lock()
			tempBlocks = []Block{}
			mutex.Unlock()
		}




		func main() {
			//env 파일 로드하기 
			err := godotenv.Load()
			if err != nil {
				log.Fatal(err)
			}
		
	
			// create genesis block
			t := time.Now()
			genesisBlock := Block{}
			genesisBlock = Block{0, t.String(), 0, calculateBlockHash(genesisBlock), "", ""}

			//디버깅을 위한 spew 사용 
			spew.Dump(genesisBlock)
			//블록체인에 genesis블록 삽입 
			Blockchain = append(Blockchain, genesisBlock)
		
	
			// start TCP and serve TCP server
			// tcp서버 시작 
			server, err := net.Listen("tcp", ":"+os.Getenv("ADDR"))
			if err != nil {
				log.Fatal(err)
			}
			//tcp 연결을 닫음 
			defer server.Close()
		
			// 고루틴을 시작한다. 
			// candidate 변수를 시작으로 tempBlocks에 candidate를 넣고 candidateBlocks 크기만큼  반복순회
			go func() {
				for candidate := range candidateBlocks {
					//go루틴 보호를 위해 동시시작하는 것을 보호하기 위해 설정 
					mutex.Lock()
					tempBlocks = append(tempBlocks, candidate)
					mutex.Unlock()
				}
			}()
		
	
			go func() {
				for {
					//승자 고르기 시작 
					pickWinner()
				}
			}()
		
	
			for {
				//func (l *TCPListener) Accept() (Conn, error): 클라이언트가 연결되면 TCP 연결(커넥션)을 리턴
				//값이 있으면 conn이고 없으면 err
				conn, err := server.Accept()
				//에러가 있을 경우 로그 출력 
				if err != nil {
					log.Fatal(err)
				}
				//커넥션 핸들러 함수 실행 
				go handleConn(conn)
			}
		}