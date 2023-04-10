package main

import (
	"embed"
	"net/http"
)

// jpgs holds the static images used on the home page.
// 추가하고 싶은 파일을 go:embed한  뒤 목록을 추가하면 빌드할 때 해당 파일들을 실행파일내에 포함시키게됨 
//go:embed static/* 
var files embed.FS // ❶ 파일을 추가합니다.

func main() {
	http.Handle("/", http.FileServer(http.FS(files))) // ❷ 파일 서버 실행
	http.ListenAndServe(":3000", nil)
}
