package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"strings"
)

// 디코더는 입력 스트림에서 JSON 값을 읽고 디코딩합니다.
/*

함수 NewDecoder(r io.Reader) *Decoder
NewDecoder는 r에서 읽는 새 디코더를 반환합니다.

디코더는 자체 버퍼링을 도입하고 요청된 JSON 값을 초과하여 r에서 데이터를 읽을 수 있습니다.
*/
// 예제 ¶
// 이 예제는 디코더를 사용하여 고유한 JSON 값 스트림을 디코딩하는 예제입니다.
func main() {
	const jsonStream = `
	{"Name": "Ed", "Text": "Knock knock."}
	{"Name": "Sam", "Text": "Who's there?"}
	{"Name": "Ed", "Text": "Go fmt."}
	{"Name": "Sam", "Text": "Go fmt who?"}
	{"Name": "Ed", "Text": "Go fmt yourself!"}
`
	type Message struct {
		Name, Text string
	}
	dec := json.NewDecoder(strings.NewReader(jsonStream))
	fmt.Println("dec", dec)
	for {
		var m Message
		if err := dec.Decode(&m); err == io.EOF {
			fmt.Println("err", err)
			break
		} else if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s: %s\n", m.Name, m.Text)
	}
}
