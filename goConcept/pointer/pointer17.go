package main

import "fmt"

type Actor struct {
	Name  string
	HP    int
	Speed float64
}

func NewActor(name string, hp int, speed float64) *Actor {

	var actor = Actor{name, hp, speed}
	return &actor

	//위 코드와 같은 그 외 방법
	//return &Actor{name, hp, speed}
	/*
		a:= new(Actor)
		a.Name = name
		a.HP = hp
		a.Speed = speed
		return a
	*/
}

func main() {
	var actor = NewActor("금토끼", 99, 100)
	fmt.Println(actor.Speed)
	fmt.Println(actor.Name)
}
