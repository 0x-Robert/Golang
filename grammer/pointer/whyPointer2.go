package main

import (
	"fmt"
	"unsafe"
)

type Data struct {
	value int
	data  [200]int
}

func ChangeData(arg *Data) { //매개변수로 Data 포인터를 받는다.
	arg.value = 999 //arg 데이터를 변경한다.
	arg.data[100] = 999
}

func main() {
	var data Data

	ChangeData(&data) //data의 메모리 주소를 인자로 넘긴다.
	fmt.Printf("value = %d\n", data.value)
	fmt.Printf("data[100] = %d\n", data.data[100]) //메모리 주소는 8바이트 숫자값이기 때문에 1608바이트의 구조체 전부가 복사되는게 아닌 8바이트만 복사된다.

	var data2 Data
	var p *Data = &data2
	fmt.Println("p", p)

	var p2 *Data = &Data{}
	fmt.Println("p2", p2)
	fmt.Println(unsafe.Sizeof(p2))
	fmt.Println("p2 size", unsafe.Sizeof(p2)+unsafe.Sizeof(p2.data))
	fmt.Println("p2 size", unsafe.Sizeof(p2.value)+unsafe.Sizeof(p2.data))
	fmt.Println("p2 size", unsafe.Sizeof(p2.value))
	fmt.Println("p2 size", unsafe.Sizeof(p2.data))

	p3 := &Data{}
	var p4 = new(Data)
	fmt.Println("p3", p3)
	fmt.Println("p4", p4)
}
