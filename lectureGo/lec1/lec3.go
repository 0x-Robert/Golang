package main

import "fmt"

// if row, err := p.tInfo.UpdateOne(ctx, filter, update); err != nil {
// 	fmt.Errorf("error updateKycStatus", "", err.Error())
// 	return err
// }

// // => row 관련 리턴은 받아도 이후 사용하지 않고, 에러 처리는 err 인자를 통해 충분하기 때문에
// // 하기와 같이 작성하는것을 추천
// if _, err := p.tInfo.UpdateOne(ctx, filter, update); err != nil {
// 	fmt.Errorf("error updateKycStatus", "", err.Error())
// 	return err
// }

func main() {
	fmt.Println("test")
}
