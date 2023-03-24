package main 


type Stringer interface {
	String() string 
}

type Student struct{

}

func main(){
	var stringer Stringer 
	stringer.(*Student ) //컴파일 에러발생 
	//Student 구조체는 String() 메서드를 포함하고 있지 않기 때문에 *Student 타입은 Stringer 인터페이스를 구현하고 있지 않습니다. 
	
}