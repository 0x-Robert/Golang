package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func barHandler(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	name := values.Get("name")
	if name == "" {
		name = "World"
	}
	id, _ := strconv.Atoi(values.Get("id"))
	fmt.Fprintf(w, "Hello %s! id:%d", name, id)
}

// mux는 multiplexer 멀티플렉서의 약자로 여러 입력 중 하나를 선택해서 반환하는 디지털 장치를 말합니다. 웹 서버는 각 URL에 해당하는 핸들러들을 등록한다음
// HTTP 요청이 왔을 때 URL에 해당하는 핸들러를 선택해서 실행하는 방식입니다. 이 핸들러를 선택하고 실행하는 구조체이름이 MUX를 제공한다고 해서 ServeMux라고 부릅니다.
// 비슷한 의미로 라우터라고 말하기도 합니다.
func main() {
	mux := http.NewServeMux() //ServeMux 인스턴스 생성
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World") //w 인스턴스에 핸들러 등록
	})

	mux.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello Bar")
	})
	mux.HandleFunc("/bar2", barHandler)
	http.ListenAndServe(":3000", mux) //mux 인스턴스 사용
}
