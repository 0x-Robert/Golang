package main

import (
	"fmt"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func main() {
	str := "peter parker. james bond"
	caser := cases.Title(language.English)
	fmt.Println("str = ", str)
	str_title := caser.String(str)
	fmt.Println("str_title_cased = ", str_title)
}
