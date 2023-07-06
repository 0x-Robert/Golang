package main

import "net/http"

func main() {
	// "/" 경로에 대한 요청이 올 때 static 폴더 아래에 있는 파일들을 제공하는 파일서버를 만듬
	//	http.Handle("/", http.FileServer(http.Dir("static"))) // static 폴더를 파일서버로 사용
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static")))) //http://localhost:3000/static/을 호출하면 이미지와 파일들을 볼수 있다. 이 중 html을 클릭!
	//하면 해당 파일로 이동해서 보여준다. 보통은 이런 정적파일은 CDN 서비스를 이용하거나 Nginx를 이용해서 서비스한다.
	http.ListenAndServe(":3000", nil) //첫번째인자는 포트와 아이피 주소, 두 번째인자는 defaultServeMux?
}
