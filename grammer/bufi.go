package main

import (
	"bufio"
	"fmt"
	"os"
)

// func NewScanner(r io.Reader) *Scanner{
// 	return &Scanner{
// 		r: r,
// 		split: ScanLines,
// 		maxTokenSize: MaxScanTokenSize,
// 	}
// }

// func main() {
//     sc := bufio.NewScanner(os.Stdin)
//     sc.Scan()	// 스트림에서 한줄 읽어옴
//     fmt.Println(sc.Text()) // 읽어온 문자열 출력
// }

func main(){
	file, err := os.OpenFile(
		"hello.txt",
		os.O_CREATE|os.O_RDWR|os.O_TRUNC, //파일이 없으면 생성, 
										  //읽기/쓰기, 파일을 연 뒤 내용 삭제 
		os.FileMode(0644),				  //파일 권한은 644 
	)
	if err != nil{ 
		fmt.Println(err)
		return 
	}
	defer file.Close() //main 함수가 끝나기 직전에 파일을 닫음 


	w := bufio.NewWriter(file) //writer인터페이스를 따르는 file로 쓰기 인스턴스  w 생성
	w.WriteString("hello world") //쓰기 인스턴스로 버퍼에 hello, world 쓰기 
	w.Flush() //버퍼의 내용을 파일에 저장 

	r := bufio.NewReader(file) 
	fi, _ := file.Stat() //파일정보 구하기 
	b := make([]byte, fi.Size()) //파일 크기만큼 바이트 슬라이스 생성 

	file.Seek(0, os.SEEK_SET) //파일 읽기 위치를 파일의 맨 처음 으로 이동
	r.Read(b) //읽기 인스턴스로 파일 내용을 읽어서 b에 저장 

	fmt.Println(string(b)) //문자열로 변환하여 b의 내용 출력 

}