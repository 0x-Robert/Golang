package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	PrintFile("hamlet.txt")
}

func PrintFile(filename string) {
	file, err := os.Open(filename) //파일 열기
	if err != nil {
		fmt.Println("파일을 찾을 수 없습니다. ", filename)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
