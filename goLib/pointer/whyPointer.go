// ch14/ex14.3/ex14.3.go
package main

import (
	"fmt"
)

type Data struct { // ❶ Data형 구조체
	value int      //8바이트
	data  [200]int //8바이트가 200개 있어서 1600바이트
	//도합1608바이트
}

func ChangeData(arg Data) { // ❷ 파라미터로 Data를 받습니다. 1608바이트가 짧은 시간에 여러번 호출되면 성능에 문제가 생길 수도 있습니다. 따라서 이런 문제를 해결할 때 포인터를 사용한다.
	arg.value = 999
	arg.data[100] = 999
}

func main() {
	var data Data

	//data와 arg는 서로 다른 메모리 공간을 갖는 변수다.
	ChangeData(data) // ❸ 인수로 data를 넣습니다.
	fmt.Printf("value = %d\n", data.value)
	fmt.Printf("data[100] = %d\n", data.data[100]) // ❹ data 필드 출력

}
