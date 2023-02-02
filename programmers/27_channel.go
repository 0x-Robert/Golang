// 채널은 파이프로, 채널 오퍼레이터 <-를 통해 값을 주고받을 수 있습니다.

// map이나 slice처럼 채널도 쓰기 전에 채널임을 선언해줘야 합니다.

// ch := make(chan int)
// ch <- v    // 채널 ch를 통해 v를 보냄.
// v := <-ch  // ch로부터 값을 전달받아, v에 할당.
// 채널은 디폴트로 상대방이 준비된 후 값을 주고받기 때문에, 별도의 동기화 과정이나 condition variable 설정 없이 goroutine을 쓸 수 있습니다.

package main

import "fmt"

// 배열 s에 모든 원소를 더해, 채널 c에게 전달하는 함수
func sum(s []int, c chan int) {
    sum := 0
    for _, v := range s {
        sum += v
    }
    c <- sum // c에게 sum을 전달
}

func main() {
    s := []int{7, 2, 8, -9, 4, 0}
    
    c := make(chan int)
    go sum(s[:len(s)/2], c)
    go sum(s[len(s)/2:], c)
    x, y := <-c, <-c // c로부터 값을 전달받음
    
	d := make(chan int)
	go sum(s[:len(s)/2], d)
    go sum(s[len(s)/2:], d)
	a, b := <-d, <-d 

    fmt.Println("sum(s[:len(s)/2], c)의 결과: x", x)
    fmt.Println("sum(s[len(s)/2:], c)의 결과: y", y)
    fmt.Println("둘의 합:", x+y)
	fmt.Println(s[:3])
	fmt.Println(s[3:])
	fmt.Println("a,b",a,b)
}

