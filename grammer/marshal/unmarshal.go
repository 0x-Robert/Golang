// // MarshalIndent is like Marshal but applies Indent to format the output.
// // Each JSON element in the output will begin on a new line beginning with prefix
// // followed by one or more copies of indent according to the indentation nesting.
// func MarshalIndent(v any, prefix, indent string) ([]byte, error) {
// 	b, err := Marshal(v)
// 	if err != nil {
// 		return nil, err
// 	}
// 	var buf bytes.Buffer
// 	err = Indent(&buf, b, prefix, indent)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return buf.Bytes(), nil
// }

//func Marshal go언어 자료형을 JSON테스트로 변환
//func MarshalIndent : go언어 자료형을 JSON 텍스트로 변환하고 사람이 보기 편하게 들여쓰기 해줌
//func Unmarshal : JSON텍스트를 go 언어 자료형으로 변환
//json형태를 go언어 자료형으로 변환
package main

import (
	"encoding/json"
	"fmt"
)

func main(){
	doc := `
	{
		
		"name" : "maria",
		"age" : 10 
	}
	`

	var data map[string]interface{} //JSON 문서의 데이터를 저장할 공간을 맵으로 선언

	json.Unmarshal([]byte(doc), &data) //doc를 바이트 슬라이스로 변환하여 넣고 data의 포인터를 넣어줌
	
	fmt.Println(data["name"], data["age"])

}