package main

func solution(my_string string, overwrite_string string, s int) string {

	answer := my_string[:s] + overwrite_string + my_string[len(overwrite_string)+s:]
	return answer
}

func main() {
	my_string := "He11oWor1d"
	overwrite_string := "lloWorl"
	s := 2
	solution(my_string, overwrite_string, s)
}
