package main

type ReadWriter interface {
	Read()
	Write()
}

type File struct {
}

func (f *File) Read() {

}

func ReadWrite(rw ReadWriter) {
	rw.Read()
	rw.Write()
}

func main() {
	//f에 File 구조체를 할당함
	//File 구조체를 리시버로 가는 Read 메서드가 있다.
	f := &File{}
	//그러나 ReadWrite 함수는 Write함수도 가지고 있는 인터페이스를 인자로 받기 때문에 인자가
	//정상적으로 동작하려면 인자는 Write()도 가지고 있어야 한다.
	ReadWrite(f)
}
