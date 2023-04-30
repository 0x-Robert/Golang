package main

import "fmt"

// PasswordError라는 사용자 정의 오류 유형을 정의합니다. 이 오류 유형은 두 개의 필드 (Len과 RequireLen)를 가집니다.
type PasswordError struct {
	Len        int
	RequireLen int
}

// PasswordError에 대한 Error() 메소드를 정의합니다. 이 메소드는 암호 길이가 짧다는 오류 메시지를 반환합니다.
func (err PasswordError) Error() string {
	return "암호 길이가 짧습니다."
}

// RegisterAccount() 함수를 정의합니다. 이 함수는 사용자 이름과 암호를 인수로 받고, 암호가 충분히 길지 않으면 PasswordError를 반환합니다.
func RegisterAccount(name, password string) error {
	if len(password) < 8 {
		return PasswordError{len(password), 8}
	}

	return nil
}

func main() {
	// err := RegisterAccount("myID", "myPw")
	err := RegisterAccount("myID", "myPwasdf")
	if err != nil {
		if errInfo, ok := err.(PasswordError); ok {
			fmt.Printf("%v Len:%d RequireLen:%d\n", errInfo, errInfo.Len, errInfo.RequireLen)
		}
	} else {
		fmt.Println("회원가입에 성공하셨습니다. ")
	}

}
