package main

import "fmt"

func merge(left [][]int, right [][]int, comparator func(item []int) int) [][]int {
	merged := make([][]int, 0)
	leftIdx := 0
	rightIdx := 0
	size := len(left) + len(right)

	for i := 0; i < size; i++ {
		if leftIdx >= len(left) {
			merged = append(merged, right[rightIdx])
			rightIdx++
		} else if rightIdx >= len(right) || comparator(left[leftIdx]) <= comparator(right[rightIdx]) {
			merged = append(merged, left[leftIdx])
			leftIdx++
		} else {
			merged = append(merged, right[rightIdx])
			rightIdx++
		}
	}
	return merged
}

func mergeSort(arr [][]int, comparator func(item []int) int) [][]int {
	var aux func(start, end int) [][]int
	aux = func(start, end int) [][]int {
		if start >= end {
			return [][]int{arr[start]}
		}
		mid := (start + end) / 2
		right := aux(start, mid)
		left := aux(mid+1, end)
		return merge(left, right, comparator)
	}
	return aux(0, len(arr)-1)
}

func shadowOfPapers(papers [][]int) int {
	coordinates := make([][]int, 0)

	for _, p := range papers {
		x, y, width, height := p[0], p[1], p[2], p[3]
		coordinates = append(coordinates, []int{x, y, y + height - 1, 1})
		coordinates = append(coordinates, []int{x + width, y, y + height - 1, -1})
	}

	comparator := func(c []int) int {
		return c[0]
	}

	sorted := mergeSort(coordinates, comparator)
	height := make([]int, 10000+1)

	for y := sorted[0][1]; y <= sorted[0][2]; y++ {
		height[y] = 1
	}

	sum := 0

	for i := 1; i < len(sorted); i++ {
		h := 0
		for _, v := range height {
			if v != 0 {
				h++
			}
		}
		x2, x1 := sorted[i][0], sorted[i-1][0]
		sum = sum + (x2-x1)*h

		y1, y2 := sorted[i][1], sorted[i][2]

		for y := y1; y <= y2; y++ {
			height[y] += sorted[i][3]
		}
	}
	return sum
}

func main() {
	result := shadowOfPapers([][]int{{0, 1, 4, 4}})
	fmt.Println(result) // --> 16

	result = shadowOfPapers([][]int{{0, 0, 4, 4}, {2, 1, 2, 6}, {1, 5, 5, 3}, {2, 2, 3, 3}})
	fmt.Println(result) // --> 36
}
