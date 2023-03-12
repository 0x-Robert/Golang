package main

import (
	"bufio"
	"os"
)

func main() {

	file, err := os.Create("file.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	// 버퍼에 데이터 쓰기
	_, err = writer.WriteString("Hello, World!")
	if err != nil {
		panic(err)
	}
	// 버퍼를 비우기 위해 Flush() 메서드 호출
	err = writer.Flush()
	if err != nil {
		panic(err)
	}

}

// 위 코드에서는 os 패키지의 Create() 메서드를 사용하여 "file.txt" 파일을 생성하고, bufio 패키지의 NewWriter() 메서드를 사용하여 파일에 버퍼링된 쓰기 작업을 위한 io.Writer 객체를 생성합니다. 그리고 writer.WriteString() 메서드를 사용하여 버퍼에 "Hello, World!" 문자열을 쓰고, 마지막으로 writer.Flush() 메서드를 호출하여 버퍼를 비워 파일에 데이터를 쓰는 작업을 수행합니다.

// 이와 같이 Go 언어에서 flush 기능은 주로 io.Writer 인터페이스를 구현한 객체의 Flush() 메서드를 호출하여 버퍼를 비우는 작업을 수행합니다.
