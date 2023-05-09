package main

import "fmt"

func main() {
	var strArray [5]string
	strArray[0] = "a"
	strArray[1] = "b"
	strArray[2] = "c"
	strArray[3] = "d"
	strArray[4] = "e"

	fmt.Println(strArray)

	var strArray2 [5]string = [5]string{"1", "2", "3", "4", "5"}
	fmt.Println(strArray2)

	strArray3 := [5]string{"10", "20", "30"}
	fmt.Println(strArray3)

	strArray4 := [...]string{"100", "200", "300"}
	fmt.Println(strArray4)
	fmt.Println(len(strArray4))

	strArray5 := [5]string{2: "codz", 4: "states"}
	fmt.Println(strArray)
	fmt.Printf("%v\n", strArray5)

}
