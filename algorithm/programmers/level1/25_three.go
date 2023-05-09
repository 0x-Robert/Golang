package main

/*
입출력 예
number	result
[-2, 3, 0, 2, -5]	2
[-3, -2, -1, 0, 1, 2, 3]	5
[-1, 1, -1, 1]	0
*/
func solution(number []int) int {
	var answer int
	for i := 0; i < len(number)-2; i++ {
		for j := i + 1; j < len(number)-1; j++ {
			for k := j + 1; k < len(number); k++ {
				if number[i]+number[j]+number[k] == 0 {
					answer++
				}
			}
		}
	}

	return answer
}

func main() {
	number := []int{-2, 3, 0, 2, -5}
	solution(number)
}
