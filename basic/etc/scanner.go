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

	for scanner.Scan() { //스캔 반복해서 단어개수를 센다. Scan()메서드가 false를 반환하면 검색을 중지 
		fmt.Println("count", count)
		count++
	}

	if err := scanner.Err(); err != nil { // 더 읽을 수 없어서 검색이 중단되면 Err 메서드는 nil을 반환한다. 
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}
	fmt.Printf("Final %d\n", count)

}
