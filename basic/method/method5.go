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
	lot := &ParkingLot{100}
	ParkCar(lot, 10)
	fmt.Println("lot", lot)

	lot.ParkCar(10)
	fmt.Println("lot", lot)
}
