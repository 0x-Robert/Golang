package main

import (
	"bufio"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func MultipleFromString(str string) (int, error) {
	scanner := bufio.NewScanner(strings.NewReader(str)) //스캐너 생성, 문자열을 한 줄씩 또는 한 단어씩 끊어 읽고자할 때 주로 사용되는 구문
	fmt.Println("scanner", scanner)
	scanner.Split(bufio.ScanWords) //한 단어씩 끊어 읽기> bufio.ScanWords, 한 줄씩 끊어 읽기 >>> bufio.ScanLines
	fmt.Println("scanner2", scanner)
	pos := 0
	a, n, err := readNextInt(scanner)
	fmt.Println("a,n,err", a, n, err)
	if err != nil {
		return 0, fmt.Errorf("Failed to readNextInt(), pos:%d err:%w", pos, err)
		//에러 감싸기
	}
	pos += n + 1
	fmt.Println("pos", pos)
	b, n, err := readNextInt(scanner)
	fmt.Println("b,n,err", b, n, err)
	if err != nil {
		return 0, fmt.Errorf("Failed to readNextInt(), pos:%d err:%w", pos, err)
	}
	fmt.Println("a * b", a*b)
	return a * b, nil

}

// 다음 단어를 읽어서 숫자로 변환하여 반환합니다.
// 변환된 숫자, 읽은 글자 수, 에러를 반환합니다.
func readNextInt(scanner *bufio.Scanner) (int, int, error) {
	if !scanner.Scan() {
		return 0, 0, fmt.Errorf("Failed to scan")
	}
	word := scanner.Text()
	fmt.Println("word", word)
	number, err := strconv.Atoi(word) //문자열을 숫자로 변환, 숫자를 문자로 바꿀때는 Itoa()함수는 숫자를 문자로 바꿈
	fmt.Println("number", number)
	if err != nil {
		return 0, 0, fmt.Errorf("Failed to convert word to int, word:%s err:%w", word, err)//%w를 사용하면 err을 감싸서 새로운 에러를 반환하게된다. 
		//에러 감싸기
	}
	fmt.Println("number len(word)", number, len(word))
	return number, len(word), nil
}

func readEq(eq string) {
	rst, err := MultipleFromString(eq)
	if err == nil {
		fmt.Println("rst", rst)
	} else {
		fmt.Println("err", err)
		var numError *strconv.NumError
		if errors.As(err, &numError) { //감싸진 에러가 NumError인지 확인. 감싸진 에러를 다시 꺼내올 때 errors.As()를 사용한다. 
			fmt.Println("NumberERror: ", numError)
		}
	}
}

func main() {
	readEq("123 3")
	readEq("123 abc")
}
