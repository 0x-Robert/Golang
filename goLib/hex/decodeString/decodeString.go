package main

import (
	"encoding/hex"
	"fmt"
	"log"
)

/*
func DecodeString
func DecodeString(s string) ([]byte, error)
DecodeString returns the bytes represented by the hexadecimal string s.

DecodeString expects that src contains only hexadecimal characters and that src has even length.
If the input is malformed, DecodeString returns the bytes decoded before the error.

DecodeString은 16진수 문자열 s로 표시되는 바이트를 반환합니다.

DecodeString은 src에 16진수 문자만 포함되고 길이가 짝수일 것으로 예상합니다.

입력이 잘못된 경우, DecodeString은 오류 이전에 디코딩된 바이트를 반환합니다.
*/
func main() {
	const s = "48656c6c6f20476f7068657221"
	decoded, err := hex.DecodeString(s)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", decoded)
}
