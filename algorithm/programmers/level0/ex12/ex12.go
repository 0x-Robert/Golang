package main

// func solution(str_list []string) []string {
// 	arr := []string{}
// 	if len(str_list) == 0 {
// 		return arr
// 	}

// 	for i, v := range str_list {

// 		if v == "l" && len(str_list) != 1 {
// 			arr = str_list[:i]

// 			break
// 		} else if v == "r" && len(str_list) != 1 {
// 			arr = str_list[i:]

// 			break
// 		} else {
// 			arr = []string{}

// 		}
// 	}
// 	return arr
// }

func solution1(str_list []string) []string {
	for idx, item := range str_list {
		if item == "l" {
			return str_list[:idx]
		} else if item == "r" {
			return str_list[idx+1:]
		}
	}
	return []string{}
}

func solution(str_list []string) []string {
	temp := make([]string, len(str_list))
	count := 0

	for i := 0; i < len(str_list); i++ {
		if str_list[i] == "l" {
			for j := 0; j < i; j++ {
				temp[count] = str_list[j]
				count += 1
			}
			break

		} else if str_list[i] == "r" {
			for j := i + 1; j < len(str_list); j++ {
				temp[count] = str_list[j]
				count += 1
			}
			break
		}
	}
	//fmt.Println(temp)
	arr := make([]string, count)
	for i := 0; i < count; i++ {
		arr[i] = temp[i]
	}
	//fmt.Println(arr)
	return arr
}

func main() {
	str_list := []string{"u", "u", "l", "r"}
	// := []string{"l"}
	solution(str_list)
}
