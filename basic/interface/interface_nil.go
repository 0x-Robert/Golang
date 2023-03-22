package main

type Attacker interface {
	Attack()
}

func main() {

	var att Attacker
	println(att)
	att.Attack()
}
