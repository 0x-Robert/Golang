package main

import "fmt"

type Data struct {
	value int      //8바이트
	data  [200]int //1600바이트
}

// func ChangeData(arg Data) >> before, 값이 안 바뀌었음
// 메모리값 수정할 수 있게 포인터 변수로 설정 >> after
// 데이터 타입의 주소를 값으로 받을 수 있다. >> *Data , after
func ChangeData(arg *Data) {
	arg.value = 999     // == *arg.value 의미적으로 표현
	arg.data[100] = 999 // == *arg.data
}

func main() {

	var a int = 10
	var b int = 20

	var p1 *int = &a
	var p2 *int = &a
	var p3 *int = &b

	fmt.Printf("p1 == p2: %v\n", p1 == p2)
	fmt.Printf("p2 == p3: %v\n", p2 == p3)

	//포인터 변수 기본 값은 nil == null
	// nil은 무효하다.

	var p *int
	if p != nil {
		//p는 유효한 메모리 주소를 가진다.
		fmt.Println("if ")

	}
	fmt.Println("p", p)

	//포인터가 주소값을 가지고 있지만 자기자신에 대한 메모리 공간도 따로 있다.
	//c는 int타입을 포인터 타입으로 변경할 수 있으나 go는 불가능함

	//포인터를 왜 쓰나??
	var data Data
	// changeData에서 아무리 값을 바꿔도 공간이 다르기 때문에
	// arg에서 데이터를 바꿔도 data는 값이 안 바뀐다.

	//ChangeData(data) before
	// *Data = Data의 타입의 주소라는 뜻 after
	ChangeData(&data)                              //메모리 주소로 호출, 주소값이니까 64비트 컴퓨터에서는 8바이트
	fmt.Printf("value=%d\n", data.value)           //8바이트
	fmt.Printf("data[100] = %d\n", data.data[100]) //1600바이트
	//왜 0이 나오는가??
	// l-value : 공간 /  r-value : 값
	// 함수인자는 무조건 우변, r-value

	//Data 타입 구조체 변수 data를 선언
	//var data Data
	//var p *Data = &data  >> data변수의 주소를 반환한다.

	//*Data타입 구조체 변수 p를 선언한다.
	//var p *Data = &Data{} >> Data 구조체를 만들어 주소를 반환한다.

	//인스턴스는 메모리에 할당된 데이터의 실체
	// 1개의 인스턴스
	// var p *Data = &Data{}
	// Data 인스턴스 하나가 만들어졌고 포인터 변수 p가 가리킨다.
	// 객체, Life cycle 개념이라고 생각하면 된다.
	// 인스턴스가 사라지면 객체가 사라진다.

	//new 내장함수, 기본값으로만 초기화 가능
	p4 := &Data{}
	var p5 = new(Data)
	fmt.Println(p4, p5)

	//인스턴스는 언제 사라지나??
	//인스턴스는 아무도 찾지 않을 때 사라진다.

}

// func TestFunc() {
// 	u := &User{}
// 	u.Age = 30 // = *u.age, u포인터 변수를 선언하고 인스턴스를 생성한다.
// 	fmt.Println(u)
// 	함수가 끝나도 인스턴스는 남아있고 다음번 가비지 컬렉터 실행시 사라진다.

// }

//정리
//인스턴스는 메모리에 생성된 데이터의 실체
//포인터를 이용해서 인스턴스를 가리키게 할 수 있음
//함수 호출 시 포인터 인수를 통해서 인스턴스를 입력받고 그 값을 변경할 수 있게 된다.
//쓸모 없어진 인스턴스는 가비지 컬렉터가 자동으로 지워줌
