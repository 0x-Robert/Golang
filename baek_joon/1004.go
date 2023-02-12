package main

//Go 언어는 io.Reader, io.Writer 인터페이스를 활용하여
//다양한 방법으로 입출력을 할 수 있습니다. 즉, 입출력 인터페이스를 받는 여러 패키지가 있으므로 화면에 출력하는 부분부터 파일, 문자열까지 같은 인터페이스로 처리할 수 있습니다.

import (
	"bufio" //buffered I/O를 구현하며 io.Reader, io.Writer를 래핑함 이 패키지를 이용하면 버퍼를 활용해 I/O의 부하를 줄일 수 있음

	"fmt"
	"os"
)

func main(){

	reader := bufio.NewReader(os.Stdin)//io.Reader 인터페이스로 io.Reader 인터페이스 따르는 읽기 인스턴스 생성
	writer := bufio.NewWriter(os.Stdout)//io.Writer 인터페이스로 io.Writer 인터페이스를 따르는 쓰기 인스턴스 생성
	//fmt.Println("reader, writer",reader, writer)
	
	fmt.Println("test")
	defer writer.Flush()


	var t int 
	fmt.Fscanln(reader, &t)

	for i := 0; i < t; i++ {
		var count int 
		var startX, startY, endX, endY, n int 
		fmt.Fscanln(reader, &startX, &startY, &endX, &endY)
		fmt.Fscanln(reader, &n)
		for j := 0; j < n; j++ {
			var planetX, planetY, r int 
			fmt.Fscanln(reader, &planetX, &planetY, &r)
			isStartContained := checkPlanetContains(startX, startY, planetX, planetY, r)
			isEndContained := checkPlanetContains(endX, endY, planetX, planetY, r)
			if isEndContained && !isStartContained || !isEndContained && isStartContained{
				count++
			}
		} 
		fmt.Fprintln(writer, count)
	}
}

func checkPlanetContains(x, y, planetX, planetY, r int) bool {
	dist := (x-planetX) * (x-planetX)+ (y-planetY)*(y-planetY)
	if dist > r*r {
		return false 
	}
	return true 
}

//입력 첫줄 : 테스트 케이스의 개수 T 
//각각 테스트 케이스에 대해 첫째 줄에 출발점과 도착점이 주어짐 (x1,y1), (x2,y2)
//두번째 줄에는 행성계의 개수 n이 주어지고 
//세번째 줄부터 n줄에 걸쳐 행성계의 중점과 반지름(cx. cy, r)
// 2 테스트 케이싀 개수 2 
// -5 1 12 1  출발점(-5,1 ) 도착점(12,1)
// 7 행성계의 개수 
// 1 1 8
// -3 -1 1
// 2 2 2
// 5 5 1
// -4 5 1
// 12 1 1
// 12 1 2 테스트케이스 1이 끝나고 행성도 7개로 끝남 

// 테스트케이스 2 시작
// -5 1 5 1 (-5,1) 에서 (5,1)이 출발점과 도착점 
// 1	행성계의 개수 1 
// 0 0 2 행성은 (0, 0 ,2) 