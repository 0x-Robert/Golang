package main

import (
	"fmt"
	"sync"
	"time"
)

// 고루틴에서 뮤텍스를 사용하지 않는 방법 중  두 번째 방법인 역할을 나누는 방법을 알아본다.
// 자동차 생산시 차체 생산 >바퀴 생산 > 도색 > 완성
// 이것은 위의 생산과정을 채널을 통해 구현하는 방법이다.
type Car struct { //Car 구조체 선언
	Body  string
	Tire  string
	Color string
}

var wg sync.WaitGroup      //대기하는 객체 생성
var startTime = time.Now() //현재 시간

func main() {
	//메인 루틴은 모든 고루틴이 종료될 때까지 대기합니다.
	tireCh := make(chan *Car)  //타이어 채널 생성
	paintCh := make(chan *Car) //도색 채널 생성

	fmt.Printf("Start Factory\n") //시작

	wg.Add(3)                       //각 공정에 1초가 걸리면 자동차 한대를 만드는데 총 3초가 걸림, 이런 시스템을 컨베이어 벨트 시스템이라고 함, 작업단위 3개 추가
	go MakeBody(tireCh)             //고루틴 생성, MakeBody()는 1초 간격으로 차체를 생성해서 tireCh 채널에 데이터를 넣어준다./ 10초 이후에 tireCh 채널을 닫아주고 루틴을 종료한다.
	go InstallTire(tireCh, paintCh) //InstallTire 루틴은 tireCh 채널에서 데이터를 읽어서 바퀴를 설치하고 paintCh 채널에 넣어준다.
	//만약 tireCh 채널이 닫히면 루틴을 종료하고 paintCh 채널을 닫아준다.
	go PaintCar(paintCh) //paintCh 채널에서 데이터를 읽어서 도색을하고 완성된 차를 출력해준다.

	wg.Wait()
	fmt.Println("Close the factory") //종료

}

func MakeBody(tireCh chan *Car) { //차체 생산
	tick := time.Tick(time.Second) //1초단위로 실행하는 tick 함수로 이해함
	fmt.Println("tick", tick)
	after := time.After(10 * time.Second)
	for {
		select {
		case <-tick:
			//make a body
			car := &Car{}           //Car 구조체 정의
			car.Body = "Sports car" //구조체 Body 설정
			tireCh <- car           //channel 데이터 입력
		case <-after: //10초 뒤 종료
			close(tireCh) //채널 닫기
			wg.Done()
			return
		}
	}
}

func InstallTire(tireCh, paintCh chan *Car) { //바퀴 설치
	for car := range tireCh {
		//make a body
		time.Sleep(time.Second)  //1초 뒤 실행
		car.Tire = "Winter tire" //구조체 tire 설정
		paintCh <- car           //channel 데이터 입력

	}
	wg.Done()
	close(paintCh) //채널 닫기
}

func PaintCar(paintCh chan *Car) { //도색
	for car := range paintCh {
		//make a body
		time.Sleep(time.Second)
		car.Color = "Red"
		duration := time.Now().Sub(startTime) //경과 시간 출력
		fmt.Printf("%.2f Complete Car: %s %s %s\n", duration.Seconds(), car.Body, car.Tire, car.Color)
	}
	wg.Done()
}
