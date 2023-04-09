package main

//Scanner를 이용해서 단어 개수를 세는 예제

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	const input = "Now is the winter of our discopntent,Made glorious summer by this sun of York. testasdfasdf adsfasd\n"
	scanner := bufio.NewScanner(strings.NewReader(input)) //스캐너 생성
	scanner.Split(bufio.ScanWords)                        //단어 단위로 검색, 띄어쓰기로 구분
	count := 0

	for scanner.Scan() { //스캔 반복
		fmt.Println("count", count)
		count++
	}

	if err := scanner.Err(); err != nil { //
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}
	fmt.Printf("Final %d\n", count)

}
