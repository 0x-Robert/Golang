package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name    string   `json:"name"`
	Age     int      `json:"age"`
	Hobbies []string `json:"hobbies"`
}

type Person2 struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func printAnything(v any) any {
	fmt.Println("any test", v)
	return v
}

func main() {
	foods := map[string]interface{}{
		"bacon": "delicious",
		"eggs": struct {
			source string
			price  float64
		}{"chicken", 1.75},
		"steak": true,
	}

	yongari := map[string]interface{}{
		"dynamic": "nft",
		"game": struct {
			fun      string
			gametime int
		}{"starcraft", 19},
		"bibim noodle": "paldo",
		"exercise":     "soccer",
	}

	p := Person{
		Name: "John",
		Age:  29,
		Hobbies: []string{
			"martial arts",
			"breakfast foods",
			"piano",
		},
	}

	newp := Person2{
		Name: "yongari",
		Age:  31,
	}

	jsonBytes, err := json.Marshal(newp)
	if err != nil {
		fmt.Println("err test")
	}

	p2 := map[string]any{}
	err = json.Unmarshal(jsonBytes, &p2)
	if err != nil {
		fmt.Println("err test")
	}

	for k, v := range p2 {
		switch c := v.(type) {
		case string:
			fmt.Printf("Item %q is a string, containing %q\n", k, c)
		case float64:
			fmt.Printf("Looks like item %q is a number, specifically %f\n", k, c)
		default:
			fmt.Printf("Not sure what type item %q is, but I think it might be %T\n", k, c)
		}
	}

	test := printAnything(0)

	fmt.Println("foods", foods)

	fmt.Println("yongari", yongari)

	fmt.Println("printAnything", test)

	fmt.Println("p", p)
}
