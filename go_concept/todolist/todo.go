package main

// go get github.com/urfave/negroni
// go get github.com/unrolled/render
import (
	"encoding/json"
	"log"
	"net/http"
	"sort"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"github.com/urfave/negroni"
)

var rd *render.Render

type Todo struct {
	ID        int    `json:"id,omitempty"`
	Name      string `json:"name"`
	Completed bool   `json:"completed,omitempty"` //json포맷으로 변환 시 항목이름이 ID가 아닌 id로 변환되고 생략가능함을 표시함
}

var todoMap map[int]Todo
var lastID int = 0

func MakeWebHandler() http.Handler {
	todoMap = make(map[int]Todo)                                             //map으로 Todo 구조체를 만듬
	mux := mux.NewRouter()                                                   //멀티플렉서 객체 생성
	mux.Handle("/", http.FileServer(http.Dir("public")))                     // 루트에서 파일서버를 만들기 >> public
	mux.HandleFunc("/todos", GetTodoListHandler).Methods("GET")              //todos get메소드
	mux.HandleFunc("/todos", PostTodoHandler).Methods("POST")                //todos post메소드
	mux.HandleFunc("/todos/{id:[0-9]}", RemoveTodoHandler).Methods("DELETE") //todos delete 메소드
	mux.HandleFunc("/todos/{id:[0-9]}", UpdateTodoHandler).Methods("PUT")    //todos put 메소드
	return mux
}

type Todos []Todo //ID로 정렬하는 인터페이스

func (t Todos) Len() int { //Todos 크기 함수
	return len(t)
}

func (t Todos) Swap(i, j int) { //Todos Swap 함수
	t[i], t[j] = t[j], t[i]
}

func (t Todos) Less(i, j int) bool { //Todos Less 함수
	return t[i].ID > t[j].ID
}

func GetTodoListHandler(w http.ResponseWriter, r *http.Request) {
	list := make(Todos, 0) //Todos를 초기화
	for _, todo := range todoMap {
		list = append(list, todo) //todoMap 크기만큼 순회하며 list에 삽입
	}
	sort.Sort(list)                 //정렬
	rd.JSON(w, http.StatusOK, list) //ID로 정렬하여 전체 목록 반환
}

func PostTodoHandler(w http.ResponseWriter, r *http.Request) {
	var todo Todo
	err := json.NewDecoder(r.Body).Decode(&todo) ///err값에 JSON 데이터 변환
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	lastID++               //lastID +1
	todo.ID = lastID       //라스트 아이디를 todo.ID에 저장
	todoMap[lastID] = todo //todoMap - lastID에 todo 객체를 설정
	rd.JSON(w, http.StatusCreated, todo)
}

// 성공 메시지 구조체
type Success struct {
	Success bool `json:"success"`
}

func RemoveTodoHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)               //배열에서 불러오기
	id, _ := strconv.Atoi(vars["id"]) //문자열을 정수로 만들기
	if _, ok := todoMap[id]; ok {
		delete(todoMap, id)
		rd.JSON(w, http.StatusOK, Success{true}) //성공메시지 반환
	} else {
		rd.JSON(w, http.StatusNotFound, Success{false}) //실패메시지 반환
	}
}

func UpdateTodoHandler(w http.ResponseWriter, r *http.Request) {
	var newTodo Todo                                //Todo 객체 선언
	err := json.NewDecoder(r.Body).Decode(&newTodo) //JSON 데이터 변환
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	if todo, ok := todoMap[id]; ok {
		todo.Name = newTodo.Name
		todo.Completed = newTodo.Completed
		rd.JSON(w, http.StatusOK, Success{true})
	}
}

func main() {

	rd = render.New()
	m := MakeWebHandler()
	n := negroni.Classic() //negroni 기본 핸들러로 감싸기
	n.UseHandler(m)

	log.Println("Started App")
	err := http.ListenAndServe(":3000", n)
	if err != nil {
		panic(err)
	}
}
