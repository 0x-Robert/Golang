package main

import (
	"fmt"
	"math/rand"
	"time"
)
func main() {
//예를 들어, rand.Intn은 0 <= n < 100 사이의 랜덤 int형 n을 반환합니다.

	fmt.Print(rand.Intn(100), ",")
	fmt.Print(rand.Intn(100))
	fmt.Println()
//rand.Float64는 0.0 <= f < 1.0의 float64형 f를 반환합니다.

	fmt.Println(rand.Float64())
//다음은 다른 범위, 예를 들면 5.0 <= f < 10.0의 랜덤 실수값을 생성하는데 사용할 수 있습니다.

	fmt.Print((rand.Float64()*5)+5, ",")
	fmt.Print((rand.Float64() * 5) + 5)
	fmt.Println()
//기본 숫자 생성기는 결정적이기 때문에, 기본적으로 매번 동일한 순서의 시퀀스를 생성합니다. 다양한 시퀀스를 생성하기 위해선, 변화하는 시드값을 주어야합니다. 참고로 이는 비밀로 하려는 랜덤수를 생성하는데에는 안전하지 않으며, 이럴땐 crypto/rand를 사용하세요.

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
// rand 패키지의 함수와 마찬가지로 rand.Rand에서 호출합니다.
//r = mrand.New(mrand.NewSource(randseed))
	fmt.Print(r1.Intn(100), ",")
	fmt.Print(r1.Intn(100))
	fmt.Println()
// 동일한 숫자의 소스를 시드값으로 하면, 동일한 시퀀스의 랜덤값이 생성됩니다.
	s2 := rand.NewSource(1)
	r2 := rand.New(s2)
	fmt.Println("s2",s2)
	//fmt.Println("r2",r2)


	fmt.Print(r2.Intn(100), ",")
	fmt.Print(r2.Intn(100))
	fmt.Println()
	s3 := rand.NewSource(42)
	r3 := rand.New(s3)
	fmt.Print(r3.Intn(100), ",")
	fmt.Print(r3.Intn(100))
}

