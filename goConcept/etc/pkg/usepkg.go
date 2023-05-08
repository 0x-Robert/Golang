package main

import (
	"fmt"
	"goproject/usepkg/custompkg"

	"github.com/guptarohit/asciigraph"
	"github.com/tuckersGo/musthaveGo/ch16/expkg"
)

// go mod tidy 명령은 go 모듈에 필요한 패키지를 찾아서 다운로드 해주고 필요한 패키지 정보를 go.mod, go.sum 파일에 적어줍니다.
// go mod tidy 명령을 실행할 때 잘못된 파일을 가져올 경우 지우고 새로 가져오는게 좋다.
func main() {
	custompkg.PrintCustom()
	expkg.PrintSample()

	data := []float64{3, 4, 5, 6, 9, 7, 5, 8, 5, 10, 2, 7, 2, 5, 6}
	graph := asciigraph.Plot(data)
	fmt.Println(graph)
}
