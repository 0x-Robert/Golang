package main

import (
	"fmt"
	"reflect"
	"strings"
)

func main() {
	//map[string]interface{}
	//go언어의 특성상 string은 key값으로 interface는 여러 값을 밸류로 쓸 수 있게 하기 위해 사용함

	data := map[string]interface{}{
		"attributes": []map[string]string{
			{"trait_type": "location", "name": "런던", "value": "뉴욕"},
			{"trait_type": "latitude", "value": "51.5072"},
			{"trait_type": "longitude", "value": "-0.1275"},
		},
		"description": "cloud",
		"image":       "https://gateway.ipfscdn.io/ipfs/QmSK8392GF11ZJ3vnpLiSvx1ciE7B5d7UT27ctkLwNa79x/goku3.png",
		"name":        "cloud",
	}

	var elements []string
	fmt.Println("elements", elements)
	fmt.Println(reflect.TypeOf(elements))
	fmt.Println("data", data)
	fmt.Println("type of", reflect.TypeOf(data))

	for k, v := range data {
		element := k + ":"
		switch v := v.(type) {
		case string:
			fmt.Println("string", element)
			element += v
		case []map[string]string:

			for _, m := range v {
				for k, s := range m {
					fmt.Println("map[string]string", element)
					element += k + ":" + s + " "
				}
			}
		}
		elements = append(elements, element)
		fmt.Println("elements", elements)
	}
	result := "[" + strings.Join(elements, ",") + "]"

	fmt.Println("result", result)

}
