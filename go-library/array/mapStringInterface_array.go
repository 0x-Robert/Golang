package main

import (
	"fmt"
	"strings"
)

func main() {
	data := map[string]interface{}{
		"attributes": []map[string]string{
			{"trait_type": "location", "name": "런던", "value": "뉴욕"},
			{"trait_type": "latitude", "value": "51.5072"},
			{"trait_type": "longitude", "value": "-0.1275"},
		},
		"description": "cloud",
		"image":       "ipfs://QmSK8392GF11ZJ3vnpLiSvx1ciE7B5d7UT27ctkLwNa79x/goku3.png",
		"name":        "cloud",
	}
	//배열을 이용한 구현
	var elements [4]string
	i := 0
	for key, value := range data {
		element := key + ":"
		fmt.Println("element", element)
		switch v := value.(type) {
		case string:
			element += v
			fmt.Println("case string : ", element)
		case []map[string]string:
			for _, m := range v {
				for k, s := range m {
					element += k + ":" + s + " "
					fmt.Println("case []map[string]string : ", element)
				}
			}
		}
		elements[i] = element
		i++
		fmt.Println("elements : ", elements)
	}
	result := "[" + strings.Join(elements[:], ",") + "]"
	fmt.Println("result", result)

}
