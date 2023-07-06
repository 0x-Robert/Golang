package calc

import "testing" // 테스트 코드는 항상 testing 패키지를 가져옴

type TestData struct {
	//테스트 데이터 구조체
	argument1 int
	argument2 int
	result    int
}

var testdata = []TestData{
	{2, 6, 8},
	{-8, 3, -5},
	{6, -6, 0},
	{0, 0, 0},
}

func TestSum(t *testing.T) { // 함수 이름은 Test로 시작,
	// // Test 다음의 함수 첫 글자는 대문자, 매개변수는 *testing.T
	// r := Sum(1, 2) // 덧셈 함수로 1과 2를 더하면 결과는 3이 나와야 함
	// if r != 3 {    // 3이 아닐 때는 테스트 실패
	// 	t.Error("결괏값이 3이 아닙니다. r=", r) // 에러 발생
	// }

	for _, d := range testdata {
		r := Sum(d.argument1, d.argument2)
		if r != d.result {
			t.Errorf(
				"%d + %d의 결과값이 %d(이)가 아닙니다. r=%d",
				d.argument1,
				d.argument2,
				d.result,
				r,
			)
		}
	}
}
