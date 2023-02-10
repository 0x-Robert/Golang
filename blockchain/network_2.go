package main

import (
	"bufio"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"io"
	"log"
	"net"
	"os"
	"strconv"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/joho/godotenv"
)

// Block represents each 'item' in the blockchain
type Block struct {
	Index     int
	Timestamp string
	BPM       int
	Hash      string
	PrevHash  string
	
}


// Blockchain is a series of validated Blocks
var Blockchain []Block
var bcServer chan []Block 


// SHA256 hashing
func calculateHash(block Block) string {
	record := string(block.Index) + block.Timestamp + string(block.BPM) + block.PrevHash
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

// create a new block using previous block's hash
func generateBlock(oldBlock Block, BPM int)(Block, error) {
    

	var newBlock Block


	t := time.Now()


	newBlock.Index = oldBlock.Index + 1
	newBlock.Timestamp = t.String()
	newBlock.BPM = BPM
	newBlock.PrevHash = oldBlock.Hash
	newBlock.Hash = calculateHash(newBlock)


	return newBlock, nil
}
// make sure block is valid by checking index, and comparing the hash of the previous block
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

// make sure the chain we're checking is longer than the current blockchain
func replaceChain(newBlocks []Block) {
	if len(newBlocks) > len(Blockchain) {
		Blockchain = newBlocks
	}
}




func main(){
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	bcServer = make(chan []Block)

	//create genesis block 
	t:= time.Now()
	genensisBlock := Block{0, t.String(), 0, "", ""}
	spew.Dump(genensisBlock)//debuging 용도
	Blockchain = append(Blockchain, genensisBlock) //블록체인 배열 변수에 제네시스 블록 삽입 


	//start TCP and serve TCP server
	server, err := net.Listen("tcp","localhost:"+os.Getenv("ADDR"))
	if err != nil{
		log.Fatal(err)
	}
	//connection이 더 이상 필요 없을 때 종료하기 위해 사용함 
	defer server.Close()


	//새로운 connection들을 받을 수 있게 무한 루프 생성 
	for {
		conn, err := server.Accept()
		if err != nil{
			log.Fatal(err)
		}
		//동시에 각 connection들을 분리해서 다룰 수 있도록 하기
		//go 루틴 go handleConn을 통해 해결
		go handleConn(conn)
	}


	
}
//한 개의 인자 (net.Conn 필요)
func handleConn(conn net.Conn){
	//각 connection 종료 가능 
	defer conn.Close()
	

	//BPM을 입력한다.
	io.WriteString(conn, "Enter a new BPM:")
    
		//새로운 스캐너를 생성합니다.
        scanner := bufio.NewScanner(conn)
    

        // take in BPM from stdin and add it to blockchain after conducting necessary validation
        go func() {
			//for scanner.Scan() 반복문은 Go 루틴안에 작성합니다. 그러면 다른 connection들과 분리하여 수행합니다.
            for scanner.Scan() {
                bpm, err := strconv.Atoi(scanner.Text())
                if err != nil {
                    log.Printf("%v not a number: %v", scanner.Text(), err)
                    continue
                }

				//이전에 생성한 gnerateBlock, isBlockValid 그리고 replaceChanin 함수들을 이용하여 데이터를 가진 블록을 생성한다.
                newBlock, err := generateBlock(Blockchain[len(Blockchain)-1], bpm)
                if err != nil {
                    log.Println(err)
                    continue
                }
                if isBlockValid(newBlock, Blockchain[len(Blockchain)-1]) {
                    newBlockchain := append(Blockchain, newBlock)
                    replaceChain(newBlockchain)
                }
				// 새로운 블록체인을 채널에 넣는 것을 의미합니다.		
				//우리가 다른 노드에 알려주기 위해 만든 네트워크에 블록체인을 넣는다.
                bcServer <- Blockchain
				//사용자가 새로운 BPM을 입력할 수 있도록 한다.
                io.WriteString(conn, "\nEnter a new BPM:")
            }
        }()

		// simulate receiving broadcast
	// simulate receiving broadcast
	go func() {
		for {
			time.Sleep(30 * time.Second)
			output, err := json.Marshal(Blockchain)
			if err != nil {
				log.Fatal(err)
			}
			io.WriteString(conn, string(output))
		}
	}()


	for _ = range bcServer {
		spew.Dump(Blockchain)
	}

}

