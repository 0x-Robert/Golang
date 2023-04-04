package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("2개 이상의 실행 인수가 필요합니다. word filepath ")
		return
	}

	word := os.Args[1]
	fmt.Println("찾으려는 단어: ", word)
	files := os.Args[2:]
	fmt.Println("files", files)
	PrintAllFiles(files)
}

func GetFileList(path string) ([]string, error) {
	//파일 경로를 넣어주면 경로에 해당하는 파일리스트를 []string 타입으로 변환
	fmt.Println("filepath.glob(path)")
	return filepath.Glob(path)
}

func PrintAllFiles(files []string) {
	for _, path := range files {
		filelist, err := GetFileList(path)
		if err != nil {
			fmt.Println("파일 경로가 잘못되었습니다. :err ", err, "path:", path)
			return
		}
		fmt.Println("찾으려는 파일 리스트")
		for _, name := range filelist {
			fmt.Println(name)
		}
	}
}
