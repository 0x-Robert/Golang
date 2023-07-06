package main

import (
	"encoding/json"
	"fmt"
)

func main(){
	data := make(map[string]interface{}) //문자열을 키로하고 모든 자료형을 저장할 수 있는 맵을 생성함
	data["name"] = "maria"
	data["age"] = 10
	
	doc, _ := json.Marshal(data) //맵을  JSON 문서로 변환
	fmt.Println(string(doc)) //문자열로 변환해서 출력

	data2 := make(map[string]interface{}) 
	data2["name"] = "yongari"
	data2["age"] = "30"
	//첫 인자는 json문서로 만들 데이터, 두번째 인자는 json 문서의 첫 칸에 표시할 문자열, 마지막은 들여쓰기할 문자
	doc2, _ := json.MarshalIndent(data2,"", " ") //맵을  JSON 문서로 변환
	fmt.Println(string(doc2))
}