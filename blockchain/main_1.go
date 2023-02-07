package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"time"
	
	"github.com/davecgh/go-spew/spew"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

//블록체인을 구성하기 위한 구조체를 정의
type Block struct {
	Index int  // Index는 블록체인에서 데이터 레코드의 위치를 나타냅니다.
	Timestamp string  // Timestamp 는 데이터가 작성될때의 시간이 작성되며 자동으로 결정됩니다.
	BPM int //pulse rate // BPM 또는 beats per minute, 이것은 pulse rate을 의미합니다.
	Hash string // Hash 는 SHA256를 이용하며 이 데이타 레코드의 식별을 하는데 사용됩니다.
	PrevHash string // PrevHash 이전 데이터 레코드의 Hash를 의미합니다.
} 

var Blockchain []Block // Blockchain 변수 설정

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

	newBlock.Index = oldBlock.Index + 1 //올드블록 인덱스에서 +1
	newBlock.Timestamp = t.String() //시간을 문자열로 
	newBlock.BPM = BPM 	//bpm 속성 설정 
	newBlock.PrevHash = oldBlock.Hash //올드블록의 해쉬 
	newBlock.Hash = calculateHash(newBlock) //hash계산하는 함수에 newBlock 인자를 넣고 호출 

	return newBlock, nil //return newBlock, newBlock 값이 없을 때 할당되는 초기값으로 생각됨

// 	- zero value는 명시적인 초기값을 할당하지 않고 변수를 만들었을 때 해당 변수가 갖게 되는 값이다.
// - nil은 포인터, 인터페이스, 맵, 슬라이스, 채널, 함수의 zero value이다. 

}
func isBlockValid(newBlock, oldBlock Block) bool {
	//올드블록 인덱스+1한 값과 뉴블록 인덱스가 같지 않으면 false를 반환
	if oldBlock.Index+1 != newBlock.Index{
		return false 
	}
	
	//올드블록 해쉬값과 뉴블록의 이전 해쉬값이 같지 않으면 false를 반환 
	if oldBlock.Hash != newBlock.PrevHash {
		return false 
	}
	
	//뉴블록을 calculateHash값에 넣어서 호출한 값과 뉴블록 해시값이 일치하지 않으면 false를 반환 
	if calculateHash(newBlock) != newBlock.Hash{
		return false 
	}
	
	//위 3가지 경우의 수를 제외한 경우에는 true를 반환하기 
	return true 
}

//블록체인 분기에 대한 함수
//새로운 블록의 길이가 기존 블록체인의 길이보다 길다면 
//블록체인 변수에 새로운 길이의 블록을 변수로 넣어주기 
func replaceChain(newBlocks []Block){
	if len(newBlocks) > len(Blockchain){
		Blockchain = newBlocks
	}
}

//gorilla/mux 패키지를 이용해서 run함수를 통해 웹서버 올리기
//.env 파일에서  8888포트를 가져옴 
func run() error {
	mux := makeMuxRouter()
	//mux := mux.NewRouter()
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

//모든 핸들러를 정의하기 위한 함수 
func makeMuxRouter() http.Handler{
	muxRouter := mux.NewRouter()
	muxRouter.HandleFunc("/", handleGetBlockchain).Methods("GET")
	muxRouter.HandleFunc("/",handleWriteBlock).Methods("POST")
	return muxRouter 
}

//get 핸들러를 통해 서버에 요청을 보내면 브라우저를 통해 블록체인을 볼 수 있음?
func handleGetBlockchain(w http.ResponseWriter, r *http.Request){
	bytes, err := json.MarshalIndent(Blockchain, "", " ")
	//에러가 nil값이 아니면 즉 에러가 있으면 다음 값들을 리턴 
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return 
	}
	//string을 쓰는 함수? 
	io.WriteString(w, string(bytes))
}

//POST request는 BPM을 담은 구조체가 필요함
//그래서 다음과 같이 구조체로 BPM을 int로 가지는 Message를 선언함
type Message struct {
	BPM int 
}

func handleWriteBlock(w http.ResponseWriter, r *http.Request){
	var m Message 
	//request body의 decord 값
	decoder := json.NewDecoder(r.Body)

	//decorder값을 디코딩했을 때 에러로 설정 이후 JSON으로 리스폰스 반환하는 값 설정 
	if err := decoder.Decode(&m); err != nil{
		respondWithJSON(w,r, http.StatusBadRequest, r.Body)
		return 
	}

	defer r.Body.Close()

	//뉴블록에 제너레이트 블록값 설정 
	newBlock, err := generateBlock(Blockchain[ len(Blockchain) -1  ], m.BPM)
	
	//에러 처리 
	if err != nil {
		respondWithJSON(w, r, http.StatusInternalServerError,m)
		return 
	}

	//블록유효성 검사, true일경우 
	if isBlockValid(newBlock, Blockchain[ len(Blockchain) -1 ]){

		//뉴블록체인은 블록체인에 뉴블록을 삽입하는 것
		newBlockchain := append(Blockchain, newBlock)
		replaceChain(newBlockchain)
		//spew.Dump는 콘솔에 찍어주기위한 함수임, 디버그를 위한 함수
		spew.Dump(Blockchain)
	}

	respondWithJSON(w,r,http.StatusCreated, newBlock)

}
func respondWithJSON(w http.ResponseWriter, r *http.Request, code int, payload interface{} ){
	//response 값에 json.
	response, err := json.MarshalIndent(payload, "", " ")
	
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("HTTP 500: Internal Server Error"))
		return 
	}
	//헤더에 코드를 삽입하고
	w.WriteHeader(code)
	//리스폰스를 삽입함
	w.Write(response)
}

func main(){
	//env에 있는 설정을 로드한다. 포트를 가져오기 
	err := godotenv.Load()
	//에러가 있으면 로그출력?
	if err != nil{
		log.Fatal(err)
	}

	go func(){
		//현재시간 생성
		t := time.Now()
		//제네시스블록 선언 
		genesisBlock := Block{0, t.String(), 0, "", ""}
		//디버깅을 위한 spew.Dump 함수실행 
		spew.Dump(genesisBlock)
		//블록체인 변수에 제네시스 변수 추가 
		Blockchain = append(Blockchain, genesisBlock)
	}()

	log.Fatal(run())
}

