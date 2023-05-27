package main

import (
	"fmt"
	"net/http"
)

func main() {
	// "/" 경로에 해당하는 HTTP 요청 수신시 호출하는 핸들러 함수를 등록한다.
	//두 번째 인수로는 실행할 함수를 입력해준다.
	// *http.Request 에는 클라이언트에서 보낸 메서드, 헤더, 바디같은 HTTP 요청 정보를 가지고 있다.
	// 여기서는 함수 리터럴을 사용해서 등록했다.
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//Fprint는 출력 스트림에 값을 쓰는 함수다.
		fmt.Fprint(w, "Hello World") //웹 핸들러 등록
	})
	//ListenAndServe를 호출해서 웹서버를 실행한다.
	//첫번째 인수로 요청을 기다릴 주소를 적어준다, 두번째 인수로는 핸들러 인스턴스를 넣어준다.
	//nil을 사용하면 DefaultServeMux를 사용하고 이 DefaultServeMux는 http.HandleFunc()함수를 호출해서 등록된 핸들러를 사용한다.
	http.ListenAndServe(":3000", nil) //웹 서버 시작
}
