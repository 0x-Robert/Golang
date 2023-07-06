package main

import "fmt"

type Actor struct {
	Name  string
	HP    int
	Speed float64
}

func NewActor(name string, hp int, speed float64) *Actor {
	return &Actor{name, hp, speed}

	//var a = Actor{name,hp,speed}
	//return &a

	// a:= new(Actor)
	// a.Name = name
	// a.HP = hp
	// a.Speed = speed
	// return a
}

func main() {
	var actor = NewActor("금토끼", 99, 100)
	fmt.Println(actor.Speed)
	fmt.Println(actor.Name)
}
