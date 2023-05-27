package main

import (
	"fmt"
	"sort"
)

type IntSlice []int //int 별칭 타입

func (p IntSlice) Len() int           { return len(p) }           //길이 반환
func (p IntSlice) Less(i, j int) bool { return p[i] < p[j] }      //비교결과 반환
func (p IntSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] } //값 바꾸기

func main() {
	s := []int{5, 2, 6, 3, 1, 4, 7, 10, 34, 254, 234, 123, 3}
	sort.Sort(sort.IntSlice(s))
	fmt.Println("s", s)
}

//천천히 가는 것을 두려워 말고 가다가 멈추는 것을 두려워하라

/*
The Ultimate Go study Guide

go ㅇ너어 오픈소스 패키지
go언어 내부 코드를 살펴보기, fmt, strings, bufio등의 코드를 직접보기
github에 있는 Go언어 오픈소스 프로젝트

*/