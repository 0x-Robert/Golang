package main

func solution2(num int) (answer string) {

	if num%2 == 0 {
		//fmt.Println("Even")
		return "Even"
	} else {
		//fmt.Println("Odd")
		return "Odd"
	}
}

func main() {
	num1 := 4
	//num2:=5
	solution2(num1)
}
