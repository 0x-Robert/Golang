package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Student struct {
	Name  string
	Age   int
	Score int
}

func MakeWebHandler() http.Handler { //핸들러 인스턴스를 생성하는 함수
	mux := http.NewServeMux()
	mux.HandleFunc("/student", StudentHandler)
	return mux
}

func StudentHandler(w http.ResponseWriter, r *http.Request) {
	var student = Student{"aaa", 16, 87}
	data, _ := json.Marshal(student) //Student 객체를 []byte로 변환
	fmt.Println("data", data)        //JSON 포맷임을 표시
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Println("w", w)
	fmt.Fprint(w, string(data)) //결과 전송
}

func main() {
	http.ListenAndServe(":3000", MakeWebHandler())
}
