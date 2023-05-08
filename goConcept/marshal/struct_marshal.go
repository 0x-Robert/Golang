package main

import (
	"encoding/json"
	"fmt"
)

/*
json:"name"과 같은 것을 struct 태그라고 합니다. 이것은 Go에서 encoding/json 패키지를 사용할 때 사용됩니다. json:"name" 태그는 해당 필드가 JSON 표현에서 name 키를 가질 것임을 나타냅니다.

즉, 구조체 필드에 json:"name" 태그가 지정되면 해당 필드가 JSON 표현에서 "name" 키와 연결됩니다. 이것은 구조체 필드 이름이 JSON 키 이름과 다를 때 유용합니다.

예를 들어, 위의 Person 구조체에서 Name 필드에 json:"name" 태그가 지정되어 있기 때문에, JSON으로 마샬링 할 때 "name" 키를 사용합니다. 예를 들어, Person{Name: "Alice", Age: 30}을 JSON으로 마샬링하면 다음과 같이 됩니다.

	{
	    "name": "Alice",
	    "age": 30,
	    "hobbies": null
	}
*/
type Person struct {
	Name    string   `json:"name"`
	Age     int      `json:"age"`
	Hobbies []string `json:"hobbies"`
}

/*
Go 언어에서 JSON을 처리할 때, encoding/json 패키지를 사용합니다.
json.Marshal() 함수는 Go 언어의 구조체나 맵 등의 데이터를 JSON 문자열로 변환하는 역할을 합니다.
json.Unmarshal() 함수는 JSON 문자열을 Go 언어의 구조체나 맵 등으로 변환하는 역할을 합니다.
*/
func main() {

	p := Person{
		Name: "John",
		Age:  29,
		Hobbies: []string{
			"martial arts",
			"breakfast foods",
			"piano",
		},
	}
	fmt.Println("p", p)

	jsonBytes, err := json.Marshal(p) //json 객체의 byte 배열임
	fmt.Println("jsonBytes", jsonBytes)
	if err != nil {
		panic(err)
	}

	jsonString := string(jsonBytes)
	fmt.Println("jsonString", jsonString)

	var p2 Person
	err = json.Unmarshal([]byte(jsonString), &p2)
	if err != nil {
		panic(err)
	}
	fmt.Println("p2", p2)
	/*
	   출력값
	   p {John 29 [martial arts breakfast foods piano]}
	   jsonBytes [123 34 110 97 109 101 34 58 34 74 111 104 110 34 44 34 97 103 101 34 58 50 57 44 34 104 111 98 98 105 101 115 34 58 91 34 109 97 114 116 105 97 108 32 97 114 116 115 34 44 34 98 114 101 97 107 102 97 115 116 32 102 111 111 100 115 34 44 34 112 105 97 110 111 34 93 125]
	   jsonString {"name":"John","age":29,"hobbies":["martial arts","breakfast foods","piano"]}
	   p2 {John 29 [martial arts breakfast foods piano]}
	*/
}
