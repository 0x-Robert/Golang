package main

import "time"

// 택배 회사 구조체를 선언하세요
type Courier struct {
	Name string //string 타입 Name을 필드로 가진다.
}

// 물품 Product 구조체를 선언하세요
type Product struct {
	Name  string //string 타입 Name
	Price int    //int타입 Price
	ID    int    //int타입 ID
}

// 소포 Parcel 구조체를 선언하세요
type Parcel struct {
	Pdt           *Product  //*Product 타입 Pdt
	ShippedTime   time.Time //time.
	DeliveredTime time.Time
}

func (c *Courier) SendProduct(pdt *Product) *Parcel {
	p := &Parcel{}
	p.Pdt = pdt
	p.ShippedTime = time.Now()
	return p
}

func (p *Parcel) Delivered() *Product {
	p.DeliveredTime = time.Now()
	return p.Pdt
}
