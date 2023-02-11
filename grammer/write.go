package main

import (
	"bytes"
	"fmt"
	"strings"
)


func main(){


	var announcements = make(chan string)
	var msg = make(chan string)
	go func(){
		announcements <- "winning validator: " + "lotteryWinner" + ""
		msg <- "announcements"
		close(announcements)
		close(msg)
	}()
	
	fmt.Println("announcements", announcements,"msg",msg)


	str1 := "Welcome"
	str2 := "Rain!"
	str := str1 + str2 
	fmt.Println(str)
	str = str1 + " my " + str2 
	fmt.Println(str)


	var b bytes.Buffer 
	b.WriteString("R")
	b.WriteString("a")
	b.WriteString("i")
	b.WriteString("n")
	fmt.Println(b.String())

	str = fmt.Sprintf("%s%s", str1, str2)
	fmt.Println(str)

	str = ""
	str += str1 
	str += str2 
	fmt.Println(str)

	mySlice := []string{"Welcome", "my", "Rain!"}
	str = strings.Join(mySlice, " * ")
	fmt.Println(str)
}