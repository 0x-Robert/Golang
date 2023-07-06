package main

import "fmt"

type Vertex struct {
	Lat, Long float64 
}

func main(){
	//map 사용 
	var mymap map[string]Vertex 

	mymap = make(map[string]Vertex)
	mymap["Bell Labs"] = Vertex{
		40.68433, -74.39967, 
	}
	fmt.Println("mymap[\"Bell Labs\"]: ", mymap["Bell Labs"])
	// map 리터럴 사용
	var mymap_literal =  map[string]Vertex{
		"Bell Labs": Vertex{
			40.1234, -74.12312, 
		}, 
		"Google": Vertex{
			37.12345,-122.325425 ,	
		},

	}
	fmt.Println("mymap_literal[\"Bell Labs\"]", mymap_literal["Bell Labs"])
	fmt.Println("mymap_literal[\"Google\"]",  mymap_literal["Google"])
}
