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
	ShippedTime   time.Time //time.Time 타입 ShippedTime 
	DeliveredTime time.Time //time.Time 타입 DeliverdTime 
}

//Courier의 메서드로 SendProduct()를 정의하세요 
func (c *Courier) SendProduct(pdt *Product) *Parcel {
	//리시버 타입은 *Courier입니다. 
	//인수로 *Product를 받습니다.
	//반환타입으로 *Parcel을 반환합니다. 
	//메서드가 호출되면 Parcel 객체를 생성하고 ShippedTime을 현재시간(time.Now()) 으로 설정
	//Parcel의 product변수는 메서드 인수로 들어온 값을 사용한다. 
	//생성된 Parcel을 반환한다. 
	p := &Parcel{}
	p.Pdt = pdt
	p.ShippedTime = time.Now()
	return p
}

//Parcel의 메서드로 Delivered()를 정의하세요 
func (p *Parcel) Delivered() *Product {
	//리시버 타입은 *Parcel입니다. 
	//인수는 없다. 
	//반환값은 *Product입니다. 
	//메서드가 호출되면 리시버의 DeliveredTime을 현재 시각으로 설정하세요
	//리시버의 Pdt를 함수 결과로 반환합니다.
	p.DeliveredTime = time.Now()
	return p.Pdt
}
