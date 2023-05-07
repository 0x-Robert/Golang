package main

import "fmt"


type Data struct {
	value int 
	data [200]int //200개 원소를 가진 int형 배열 
}

func ChangeData( d Data){
	d.value = 999;
	d.data[100] = 999; 
}

func main(){
	var data Data 

	//함수 인자로 전달되면서 함수 스코프 내의 지역변수로 값 복사 
	//struct 크기는 201 * 8 바이트만큼 할당 
	//함수호출할때마다 복사, 
	ChangeData(data)
	fmt.Println(data.value)
	fmt.Println(data.data[100])

}