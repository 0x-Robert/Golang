package level1

func maxNumber(array []int) int {
	max := array[0]
	for i := 0; i < len(array); i++ {
		if array[i] > max {
			max = array[i]
		}
	}
	return max
}

func minNumber(array []int) int {
	min := array[0]
	for i := 0; i < len(array); i++ {
		if array[i] < min {
			min = array[i]
		}
	}
	return min
}

func commonCheck(arr1 []int, arr2 []int) []int {
	var common []int
	for _, v1 := range arr1 {
		for _, v2 := range arr2 {
			if v1 == v2 {
				common = append(common, v1)
			}
		}
	}
	return common
}

func solution22(n int, m int) []int {

	var max int
	var answer []int
	arr1 := make([]int, 0)
	arr2 := make([]int, 0)
	arr3 := make([]int, 0)
	arr4 := make([]int, 0)

	if n > m {
		max = n
	} else {
		max = m
	}

	for i := 1; i <= max; i++ {
		if m%i == 0 {
			arr1 = append(arr1, i)
		}
		if n%i == 0 {
			arr2 = append(arr2, i)
		}
		arr3 = append(arr3, n*i)
		arr4 = append(arr4, m*i)
	}

	comDenominator := commonCheck(arr1, arr2)
	comMultiple := commonCheck(arr3, arr4)
	gcd := maxNumber(comDenominator)
	lcm := minNumber(comMultiple)
	answer = append(answer, gcd)
	answer = append(answer, lcm)

	return answer
}

func gcd_lcm_moosik() {

	// n := 3
	// m := 12
	n := 2
	m := 5
	solution22(n, m)
}
