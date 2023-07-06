package main

import "fmt"

func solution(id_pw []string, db [][]string) string {

	id := id_pw[0]
	pw := id_pw[1]
	//fmt.Println("id,pw", id, pw)

	answer := ""
	for _, v := range db {
		//fmt.Println(i, v[0], v[1])

		if id == v[0] && pw != v[1] {
			answer = "wrong pw"
		} else if id == v[0] && pw == v[1] {
			answer = "login"
		} else if id != v[0] && pw != v[1] {
			answer = "fail"
		}
	}
	fmt.Println(answer)
	return answer
}

func main() {
	db := [][]string{
		{"rardss", "123"}, {"yyoom", "1234"}, {"meosseugi", "1234"},
	}
	id_pw := []string{"meosseugi", "1234"}
	solution(id_pw, db)
}
