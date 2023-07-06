package main

import (
	"log"
	"os"
)

func main() {
	f, err := os.OpenFile("notes.txt", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		log.Fatal(err)
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
		/*
			PathError의 포인터타입을 반환하는데 PathError 타입에 별도 패키지가 명시되어 있지 않으므로
			PathError 타입은 Open()함수가 속한 패키지와 같은 os 패키지에 속합니다. 그래서 *os.PathError이 답이 됩니다.
		*/
		//에러가 발생하면 *os.PathError에러 타입을 반환한다.
	}
}
