package main

import (
	"encoding/csv"
	"fmt"
	"strings"
)

func main() {
	csvString := "John,Doe,120 jefferson st.,Riverside, NJ, 08075\n" +
		"Jack,McGinnis,220 hobo Av.,Phila, PA,09119\n" +
		"\"John \"\"Da Man\"\"\",Repici,120 Jefferson St.,Riverside, NJ,08075\n" +
		"Stephen,Tyler,\"7452 Terrace \"\"At the Plaza\"\" road\",SomeTown,SD, 91234\n" +
		",Blankman,,SomeTown, SD, 00298\n" +
		"\"Joan \"\"the bone\"\", Anne\",Jet,\"9th, at Terrace plc\",Desert City,CO,00123"

	r := csv.NewReader(strings.NewReader(csvString))
	fmt.Println("r", r)
	records, err := r.ReadAll()
	if err != nil {
		panic(err)
	}

	fmt.Println("records", records)
}

/*
실행 결과

r &{44 0 0 false false false false 0xc000062180 0 0 [] [] [] [] []}
records [[John Doe 120 jefferson st. Riverside  NJ  08075] [Jack McGinnis 220 hobo Av. Phila  PA 09119] [John "Da Man" Repici 120 Jefferson St. Riverside  NJ 08075] [Stephen Tyler 7452 Terrace "At the Plaza" road SomeTown SD  91234] [ Blankman  SomeTown  SD  00298] [Joan "the bone", Anne Jet 9th, at Terrace plc Desert City CO 00123]]
*/
