package main 

type Reader interface{
	Read()
}

type Close interface{
	Close()
}

type File struct{

}

type (f *File) Read(){

}

func ReadFile(reader Reader){
	//Reader 인터페이스 변수를 Closer인터페이스로 타입 변환합니다.
	//런 타임 에러가 발생한다. 
	c := reader.(Closer)
	c.Closer()
}

func main(){
	file := &File{}
	ReadFile(file)
}