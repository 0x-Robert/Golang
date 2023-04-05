package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath" //파일 경로를 탐색할 때 쓰는 패키지
)

func main() {

	//go run file_search_ext.go /home/robertseo/work/golang/basic/file_search .txt
	// 1번째 인자 = 경로, 2번째 인자 : 확장자

	//바이너리 버전 : ./file_search_ext /home/robertseo/work/golang/basic/file_search .go

	//일치하는 파일은 파일 슬라이스에 저장됩니다.
	var files []string

	//입력받은 경로를 저장하고 이 경로에서 검색합니다.
	path_info := os.Args[1]

	//2번째 입력받은 문자는 확장자로 .go, .txt로 입력한다.
	ext := os.Args[2]
	// fmt.Println("ext", ext)
	// fmt.Println(reflect.TypeOf(ext))

	//filepath.Walk의 첫 번째 매개변수는 검색하고자 하는 디렉터리입니다. 두 번째 매개변수는 각 파일 또는 디렉터리를 방문하기 위해 filepath.Walk가 호출하는 함수인 walk 함수입니다.
	err := filepath.Walk(path_info, func(path string, info os.FileInfo, err error) error {

		//에러로그 출력
		if err != nil {
			fmt.Println(err)
			return nil
		}

		/*
			info.IsDir()은 파일 정보를 나타내는 구조체 info의 메서드 중 하나입니다. 이 메서드는 파일이 디렉토리인지 여부를 나타내는 부울값을 반환합니다. !info.IsDir()는 이 값이 false일 경우, 즉 파일이 디렉토리가 아닐 경우
			filepath.Ext(path)는 주어진 파일 경로 path의 확장자를 반환하는 함수입니다. 이 확장자는 경로의 마지막 점(.) 이후의 문자열로 결정됩니다.
		*/
		//fmt.Println("ext", ext)
		if !info.IsDir() && filepath.Ext(path) == ext {
			files = append(files, path)
		}

		return nil
	})
	//에러 출력 부분
	if err != nil {
		log.Fatal(err)
	}

	//files를 순회하면서 찾은 파일 정보를 한줄씩 출력해준다.
	for _, file := range files {
		fmt.Println(file)
	}
}
