package main 

import 
("fmt"
"strconv" )

func solution(n int64) []int {
    
	arr := []int{}
	str := strconv.Itoa(int(n))
	runes := []rune(str)
	
	for i, j := 0, len(runes)-1; i<j; i, j = i+1, j-1{
		runes[i], runes[j] = runes[j], runes[i]
	}
 	reverseStr2  :=  string(runes)
	for _, r := range reverseStr2{
		num,_ := strconv.Atoi(string(r))
		arr = append(arr,num)
	}
	
	

	 return arr
}



func solution2(n int64) []int {
   
	var answer []int 

	temp := n 
	for temp > 0 {
		answer = append(answer, int(temp%10))
		fmt.Println(answer)
		temp /=10 //남은 자릿수만큼 저장 
		fmt.Println(temp)
	}
	return answer 
}

func main(){
	num := int64(12345)
	solution2(num)
}