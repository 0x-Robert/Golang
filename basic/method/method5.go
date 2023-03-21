package main

import "fmt"

type ParkingLot struct {
	LotSize int
}

// 함수
func ParkCar(lot *ParkingLot, carSize int) {
	lot.LotSize -= carSize
}

// func과 함수명 사이에 리시버가 있기 때문에 메서드
func (p *ParkingLot) ParkCar(carSize int) {
	p.LotSize -= carSize
}


func main() {
	//메모리 주소값을 선언하면서 구조체에 100을 할당해서 lot 변수에 담는다. 
	lot := &ParkingLot{100}
	ParkCar(lot, 10)
	fmt.Println("lot", lot)

	lot.ParkCar(10)
	fmt.Println("lot", lot)
}
