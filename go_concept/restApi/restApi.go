// ch30/ex30.3/ex30.3.go
package main

import (
	"encoding/json"
	"net/http"
	"sort"
	"strconv"

	"github.com/gorilla/mux"
)

// RestAPI는 여러가지 분산웹 아키텍처를 합친 개념이다. 특징으로는 자기표현적인 URL, 메서드를 통한 행동표현, 무상태, 캐시활용에 유용하다
// 서버렌더링 특징 : 요청이 많아지면 웹서버 응답이 느려짐, 서버구조가 단순, URLㅏㄴ 보고서는 어떤 요청인지 알 수 없다.  캐시사용이 어렵다, 여러 서버로 분산하기 어렵다.
type Student struct {
	Id    int
	Name  string
	Age   int
	Score int
}

var students map[int]Student // ❶ 학생 목록을 저장하는 맵
var lastId int

func MakeWebHandler() http.Handler {
	mux := mux.NewRouter() // ❷ gorilla/mux를 만듭니다.
	mux.HandleFunc("/students", GetStudentListHandler).Methods("GET")
	//-- ❸ 여기에 새로운 핸들러 등록 --//
	mux.HandleFunc("/students/{id:[0-9]+}", GetStudentHandler).Methods("GET")
	mux.HandleFunc("/students", PostStudentHandler).Methods("POST")
	mux.HandleFunc("/students/{id:[0-9]+}", DeleteStudentHandler).Methods("DELETE")

	students = make(map[int]Student) // ❹ 임시 데이터 생성
	students[1] = Student{1, "aaa", 16, 87}
	students[2] = Student{2, "bbb", 18, 98}
	lastId = 2

	return mux
}

type Students []Student // Id로 정렬하는 인터페이스
func (s Students) Len() int {
	return len(s)
}
func (s Students) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s Students) Less(i, j int) bool {
	return s[i].Id < s[j].Id
}

func GetStudentListHandler(w http.ResponseWriter, r *http.Request) {
	list := make(Students, 0) // ➎ 학생 목록을 Id로 정렬
	for _, student := range students {
		list = append(list, student)
	}

	sort.Sort(list)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(list) // ➏ JSON 포맷으로 변경
}

// //*
// mux.Vars는 Go 언어의 gorilla/mux 패키지에서 제공하는 변수 맵입니다. mux.Vars 맵은 http.Request 객체와 연결된 URL 경로 변수의 값을 저장합니다.

// 예를 들어, /{category}/{id}와 같은 URL 패턴이 있다고 가정해보겠습니다. 이 패턴을 gorilla/mux로 처리하면, http.Request 객체의 URL 필드에서 Path 필드로 해당 경로를 추출 할 수 있습니다. Path 필드는 /category/id와 같은 형식의 문자열입니다. 그러나 우리가 관심을 가지는 것은 이러한 경로에서 추출한 category와 id 변수 값입니다.

// 이러한 경로 변수 값을 추출하려면 mux.Vars 맵을 사용합니다. mux.Vars 맵은 http.Request 객체와 연결된 category와 id 변수 값을 저장합니다. 이 맵은 map[string]string 타입으로 선언되어 있으며, 변수 이름이 키이고 변수 값이 값입니다.

// 예를 들어, 다음과 같은 코드를 사용하여 category와 id 변수 값을 추출 할 수 있습니다.

// go
// Copy code

// 	func handlerFunc(w http.ResponseWriter, r *http.Request) {
// 	    vars := mux.Vars(r)
// 	    category := vars["category"]
// 	    id := vars["id"]
// 	    // category와 id 변수 값을 사용하여 처리
// 	}
// */

func GetStudentHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) // ❶ id를 가져옵니다.
	id, _ := strconv.Atoi(vars["id"])
	student, ok := students[id]
	if !ok {
		w.WriteHeader(http.StatusNotFound) // ❷ id에 해당하는 학생이 없으면 에러
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(student)
}

func PostStudentHandler(w http.ResponseWriter, r *http.Request) {
	var student Student
	err := json.NewDecoder(r.Body).Decode(&student) // ❶ JSON 데이터 변환
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	lastId++ // ❷ id를 증가시킨후 맵에 등록
	student.Id = lastId
	students[lastId] = student
	w.WriteHeader(http.StatusCreated)
}

func DeleteStudentHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) // ❶ id를 가져옵니다.
	id, _ := strconv.Atoi(vars["id"])
	_, ok := students[id]
	if !ok {
		w.WriteHeader(http.StatusNotFound) // ❷ id에 해당하는 학생이 없으면 에러
		return
	}
	delete(students, id)
	w.WriteHeader(http.StatusOK) // ❸ StatusOK 반환
}

func main() {
	http.ListenAndServe(":3000", MakeWebHandler())
}
