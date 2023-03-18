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
		stdin.ReadString('\n')
	}
	return n, err
}

func main() {

	rand.Seed(time.Now().UnixNano())
	leftNumber := 1000

	for {

		randN := rand.Intn(5) + 1
		//fmt.Println("randN", randN)

		inputNum, err := InputIntValue()
		//fmt.Println("inputValue", inputNum)

		if err != nil {
			fmt.Println("숫자만 입력하세요 ")

		} else if inputNum > 5 || inputNum < 1 {

			fmt.Println("1부터 5사이만 입력하세요")
		} else {
			fmt.Println("else")
			if inputNum == randN {
				leftNumber = leftNumber + 500
				fmt.Println("숫자를 맞추셨습니다. 축하드립니다.")
				if leftNumber >= 5000 {
					fmt.Println("게임을 승리하셨습니다.")
					break
				}

			} else {
				leftNumber = leftNumber - 100
				fmt.Println("숫자를 맞추지 못하셨습니다. 아쉽네요")
				if leftNumber <= 0 {
					fmt.Println("게임에서 패배하셨습니다.")
					break
				}
			}
		}
		fmt.Println("leftnumber", leftNumber)

	}

}
