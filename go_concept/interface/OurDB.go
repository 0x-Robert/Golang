package AwesomeDB

type OurDB struct {
	Name string
}

func (db *OurDB) GetData() string {

}

func (db *OurDB) WriteData(data string) {

}

func (db *OurDB) Close() error {

}

//위의 OurDB 구조체의 모든 공개된 메서드를 이용하는 인터페이스를 만들어보라

type AllOurDB interface {
	GetData() string
	WRiteData(data string)
	Close() error
}
