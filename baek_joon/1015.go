package main

import (
	"fmt"
	"sort"
)


func main(){

	// 수도코드
	// 첫 번째 줄 입력 받아서 n에 저장
	// 공백 기준으로 배열에 전부 삽입
	// 이후 배열을 복사하고 
	// 복사한 배열에서 정렬한 뒤 딕셔너리 같이 키와 밸류를 만듬
	// 이후 복사한 배열의 딕셔너리의 키를 출력함 

	var n int 
	fmt.Scanln(&n)
	
	//a를 배열로 초기화 후 입력받는 변수를 a에 담기 
	a := make([]int, n)
	for i := 0; i < n; i++{
		fmt.Scan(&a[i])
	}

	//p를 초기화 하고 p의 배열을 만들되 n의 숫자만큼 순차적으로 만듬 
	p := make([]int, n)
	for i := 0; i < n; i++{
		p[i]=i 
	}
	//fmt.Println("before a",a)
	//fmt.Println("before p",p)
	//2, 3, 1을 밸류로하고 0, 1, 2를 키로 한 뒤 
	//정렬했을 때 인덱스는 2, 0, 1이된다. 
	//다음은 a정렬을 기준으로 p로 오름차순 정렬했을 때 p의 키값? 인덱스로 정렬된 값이다. 
	sort.Slice(p, func(i, j int) bool { return a[p[i]] < a[p[j]]})
	//2, 3, 1 >>>>  2:0 / 3:1 / 1:2
	// 2 > 0 > 1  위의 밸류값으로 다시 인덱스를 잡으면 다음과 같음 
	

	//fmt.Println("after a",a)
	//fmt.Println("after p",p)

	//이후 p의 인덱스에 있는대로 b배열을 만들면 원래 배열의 인덱스별로 오름차순이 가능함
	b := make([]int, n)
	for i := 0; i < n; i++{
		b[p[i]] = i 
		//b[p[0]]=0 , b[2] = 0 
		//b[p[1]]=1 , b[0] = 1 
		//b[p[2]]=2 , b[1] = 2 
	}

	//다음과 같이 1 2 0 으로 출력이 가능함
	for i := 0; i < n; i++{
		fmt.Print(b[i], " ")
		
	}

	
}