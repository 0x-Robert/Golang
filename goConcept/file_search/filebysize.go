package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

// 경로랑 파일 확장자 입력하면, 경로에서 해당 파일 확장자 전부 찾아내기

var files []string

func VisitFile(path string, info os.FileInfo, err error) error {

	if err != nil {
		fmt.Println(err)
		return nil
	}

	file_size := info.Size()

	if !info.IsDir() && file_size > 1024*1024 {
		files = append(files, path)
	}

	return nil
}

func main() {
	root := os.Args[1]
	err := filepath.Walk(root, VisitFile)

	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Println("file", file)
	}

}
