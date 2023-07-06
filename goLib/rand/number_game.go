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
	answer := rand.Intn(100)
	cnt := 1

	fmt.Println("a", answer)
	for {

		fmt.Printf("숫자값을 입력하세요.")
		n, err := InputIntValue()
		if err != nil {
			fmt.Println("숫자만 입력하세요")
		} else {

			if answer > n {

				fmt.Println("입력한 숫자보다 큽니다. ")

			} else if answer < n {

				fmt.Println("입력한 숫자보다 작습니다.")

			} else {

				fmt.Println("숫자를 맞추셨습니다.", "시도한 숫자 : ", cnt)
				break
			}
			cnt++
		}

	}
}
