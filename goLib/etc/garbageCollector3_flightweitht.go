// appB/exB5/exB5.go
package main

import "fmt"

//메모리 쓰레기를 재활용하는 방법
//자주할당되는 객체를 객체 풀에 넣었다가 다시꺼내쓰기
//플라이웨이트 패턴방식

func main() {
	fac := NewFlyweightFactory(1000) // ❶ 객체 공장
	for i := 0; i < 1000; i++ {
		obj := fac.Create() // ❷ 1000번 만들고 버린다
		obj.Somedata = "Somedata"
		fac.Dispose(obj)
	}

	fmt.Println("AllocCnt:", fac.AllocCnt)
}

type FlyweightFactory struct {
	pool     []*Flyweight
	AllocCnt int
}

func (fac *FlyweightFactory) Create() *Flyweight {
	var obj *Flyweight
	if len(fac.pool) > 0 { // ❸ 재활용
		obj, fac.pool = fac.pool[len(fac.pool)-1], fac.pool[:len(fac.pool)-1]
		obj.Reuse()
	} else {
		obj = &Flyweight{} // ❹ 새로 만든다
		fac.AllocCnt++
	}
	return obj
}

func (fac *FlyweightFactory) Dispose(obj *Flyweight) { // ➎ 반환
	obj.Dispose()
	fac.pool = append(fac.pool, obj)
}

func NewFlyweightFactory(initSize int) *FlyweightFactory {
	return &FlyweightFactory{pool: make([]*Flyweight, 0, initSize)}
}

type Flyweight struct {
	Somedata   string
	isDisposed bool
}

func (f *Flyweight) Reuse() {
	f.isDisposed = false
}

func (f *Flyweight) Dispose() {
	f.isDisposed = true
}
