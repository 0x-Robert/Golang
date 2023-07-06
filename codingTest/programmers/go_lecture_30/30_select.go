// select문을 쓰면 goroutine은 적어도 하나의 case가 실행될 수 있을 때까지 블락됩니다.

// case가 준비되면 case를 실행하는데, 여러 case가 준비된 경우는 준비된 case 중 하나를 무작위로 선택해 실행합니다. 준비된 케이스가 없을 때에는 select내의 default 케이스가 실행되며, 블락 없이 값을 주거나 받을 때 사용합니다.

// select {
// case i := <-c:
//     // i를 사용
// default:
//     // c로부터 받는 게 블락
// }

package main

import (
	"fmt"
	"time"
)

func main(){
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Microsecond)
	for {
		select{
		case <-tick:
			fmt.Println("100\t밀리초 지났어요")
		case <-boom:
			fmt.Println("500\t밀리초 지났어요. 종료합니다.!")
		default:
			fmt.Println("50\t밀리초 지났어요.")
			time.Sleep(50 * time.Microsecond)
		}
		
	}
}