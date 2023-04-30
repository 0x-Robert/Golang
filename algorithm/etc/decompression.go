package main

import "fmt"

func decompression(image [][]int) string {
	//재귀함수
	// rs=row start / re=row end / cs=col start / ce=col end
	var aux func(int, int, int, int, [][]int) string
	aux = func(rs, re, cs, ce int, image [][]int) string {
		//base case
		if rs == re {
			return fmt.Sprintf("%d", image[rs][cs])
		}

		//4분면으로 사각형을 나눈다.
		midRow := (rs + re) / 2
		midCol := (cs + ce) / 2
		leftUpper := aux(rs, midRow, cs, midCol, image)
		rightUpper := aux(rs, midRow, midCol+1, ce, image)
		leftDown := aux(midRow+1, re, cs, midCol, image)
		rightDown := aux(midRow+1, re, midCol+1, ce, image)

		//4분면을 조합하기
		result := leftUpper + rightUpper + leftDown + rightDown
		if result == "1111" {
			return "1"
		} else if result == "0000" {
			return "0"
		} else {
			return "X" + result
		}
	}
	return aux(0, len(image)-1, 0, len(image)-1, image)
}

func main() {
	image := [][]int{
		{1, 0, 1, 1},
		{0, 1, 1, 1},
		{0, 0, 1, 1},
		{0, 0, 0, 0},
	}
	result := decompression(image)
	fmt.Println(result)
}
