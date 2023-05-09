package main

import (
	"bufio"
	"fmt"
	"os"
)

//기본적으로 os 패키지의 파일 입출력은 버퍼가 없는 저수준 입출력 입니다. 이는 큰 용량의 파일을 처리할 때 큰 부하를 일으킵니다.

//bufio 를 사용하면 버퍼를 활용해 하드웨어 부담을 줄일 수 있습니다.

func main() {
	file1, _ := os.Create("./res.txt")
	defer file1.Close()
	fmt.Fprint(file1, "Hello World")

	f, _ := os.Create("./res2.txt")
	w := bufio.NewWriter(f)
	fmt.Fprint(f, "Hello yongari")
	fmt.Fprint(w, "W")
}
