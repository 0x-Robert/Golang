package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func main(){
	data := "hello yongari"
	hash := sha256.New()

	hash.Write([]byte(data))

	md := hash.Sum(nil)
	mdStr := hex.EncodeToString((md))

	fmt.Println(mdStr)

}