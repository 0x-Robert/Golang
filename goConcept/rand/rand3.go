package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

var stdin = bufio.NewReader(os.Stdin)

func InputIntValue() (int, error) {
	var n int
	_, err := fmt.Scanln(&n)
	if err != nil {
		// 입력 버퍼를 비우기 위해 bufio.Scanner를 사용함
		scanner := bufio.NewScanner(stdin)
		if scanner.Scan() {
			err = scanner.Err()
		}
	}
	return n, err
}

func main() {
	rand.Seed(time.Now().UnixNano())
	leftNumber := 1000

	for {
		randN := rand.Intn(5) + 1
		fmt.Println("randN", randN)

		inputNum, err := InputIntValue()
		fmt.Println("inputValue", inputNum)

		if err != nil {
			fmt.Println("숫자만 입력하세요 ")
		} else if inputNum < 1 || inputNum > 5 {
			fmt.Println("1부터 5사이만 입력하세요")
		} else {
			fmt.Println("else")
			if inputNum == randN {
				leftNumber += 500
				fmt.Println("숫자를 맞추셨습니다. 축하드립니다.")
			} else {
				leftNumber -= 100
				fmt.Println("숫자를 맞추지 못하셨습니다. 아쉽네요")
			}
			// leftNumber		// leftNumber의 값을 0과 5000 사이로 제한함
			if leftNumber < 0 {
				leftNumber = 0
			} else if leftNumber > 5000 {
				leftNumber = 5000
			}
		}

		fmt.Println("leftnumber", leftNumber)

		if leftNumber >= 5000 || leftNumber <= 0 {
			fmt.Println("게임이 종료됐습니다.")
			break
		}
	}
}
