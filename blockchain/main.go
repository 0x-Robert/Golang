package main

import (
	"crypto/sha256"
	"encoding/hex"
	"log"
	"net/http"
	"os"
	"time"
	//     "github.com/davecgh/go-spew/spew"
	//     "github.com/gorilla/mux"
	//     "github.com/joho/godotenv"
)

// import (
//     "crypto/sha256"
//     "encoding/hex"
//     "encoding/json"
//     "io"
//     "log"
//     "net/http"
//     "os"
//     "time"
//     "github.com/davecgh/go-spew/spew"
//     "github.com/gorilla/mux"
//     "github.com/joho/godotenv"
// )

//블록체인을 구성하기 위한 구조체를 정의
type Block struct {
	Index int  // Index는 블록체인에서 데이터 레코드의 위치를 나타냅니다.
	Timestamp string  // Timestamp 는 데이터가 작성될때의 시간이 작성되며 자동으로 결정됩니다.
	BPM int //pulse rate // BPM 또는 beats per minute, 이것은 pulse rate을 의미합니다.
	Hash string // Hash 는 SHA256를 이용하며 이 데이타 레코드의 식별을 하는데 사용됩니다.
	PrevHash string // PrevHash 이전 데이터 레코드의 Hash를 의미합니다.
} 

var Blockchain []Block // Blockchain 변수명에 배열을 선언?

//블록에 대한 SHA256 해시를 생성하는 함수 만들기
func calculateHash(block Block) string {
	record := string(block.Index) + block.Timestamp + string(block.BPM) + block.PrevHash //record 변수에 블록 인덱스+ 블록.timestamp + 블록.BPM + 블록.prevhash(앞에있는 해쉬) 연결해서 만듬 
	h := sha256.New() //해시함수를 새로 생성
	h.Write([]byte(record)) //레코드를 SHA256으로 해시화한 후 문자열 형태의 해시 값으로 반환함
	hashed := h.Sum(nil) //그리고 h값을 sum으로 전부 더한다? 
	return hex.EncodeToString(hashed) // hashed 최종값을 encode해서 hex값으로 리턴한다?
}

//블록 생성하는 함수
func generateBlock(oldBlock Block, BPM int)(Block, error){
	var newBlock Block 

	t := time.Now()

	newBlock.Index = oldBlock.Index + 1
	newBlock.Timestamp = t.String()
	newBlock.BPM = BPM 
	newBlock.PrevHash = oldBlock.Hash 
	newBlock.Hash = calculateHash(newBlock)

	return newBlock, nil 

}
func isBlockValid(newBlock, oldBlock Block) bool {
	if oldBlock.Index+1 != newBlock.Index{
		return false 
	}
	if oldBlock.Hash != newBlock.PrevHash {
		return false 
	}
	if calculateHash(newBlock) != newBlock.Hash{
		return false 
	}
	return true 
}


func replaceChain(newBlocks []Block){
	if len(newBlocks) > len(Blockchain){
		Blockchain = newBlocks
	}
}

func run() error {
	mux := makeMuxRouter()
	httpAddr := os.Getenv("ADDR")
	log.Println("Listening on ", os.Getenv("ADDR"))
	s:= &http.Server{
		Addr: ":" + httpAddr,
		Handler: mux,
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	if err := s.ListenAndServe(); err != nil{
		return err 
	}
	return nil 
}