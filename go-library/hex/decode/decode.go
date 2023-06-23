package main

// func Decode
// func Decode(dst, src []byte) (int, error)
// Decode decodes src into DecodedLen(len(src)) bytes, returning the actual number of bytes written to dst.

// Decode expects that src contains only hexadecimal characters and that src has even length. If the input is malformed, Decode returns the number of bytes decoded before the error.

/*

디코딩은 src를 디코딩된 길이(len(src)) 
바이트로 디코딩하여 dst에 쓰여진 실제 바이트 수를 반환합니다.

Decode는 src에 16진수 문자만 포함되고 길이가 짝수일 것으로 예상합니다. 
입력이 잘못된 경우, Decode는 오류 이전에 디코딩된 바이트 수를 반환합니다.

*/
import (
	"encoding/hex"
	"fmt"
	"log"
)

func main() {
	src := []byte("48656c6c6f20476f7068657221")

	dst := make([]byte, hex.DecodedLen(len(src)))
	n, err := hex.Decode(dst, src)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", dst[:n])
}
