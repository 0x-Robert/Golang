package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// 찾은 라인 정보
type LineInfo struct { //찾은 결과 정보
	lineNo int
	line   string
}

// 파일 내 라인 정보
type FindInfo struct {
	filename string
	lines    []LineInfo
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("2개 이상의 실행 인수가 필요합니다. file_search_done yong filepath")
		return
	}

	word := os.Args[1] //찾으려는 단어
	files := os.Args[2:]
	//findInfos := []FindInfo{}
	for _, path := range files {
		//파일 찾기
		findInfos = append(findInfos, FindWordInAllFiles(word, path)...)
	}
	for _, findInfo := range findInfos {
		fmt.Println(findInfo.filename)
		fmt.Println("---------------")
		for _, lineInfo := range findInfo.lines {
			fmt.Println("\t", lineInfo.lineNo, "\t", lineInfo.line)
		}
		fmt.Println("---------------")
		fmt.Println()
	}

}

func GetFileList(pattern string) ([]string, error) {

	filelist := []string{}
	err = filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			matched, _ := filepath.Match(pattern, info.Name())
			if matched {
				filelist = append(filelist, path)
			}
		}
		return nil
	})

	if err != nil {
		return []string{}, err
	}
	return filelist, nil, err
}

func FindWordInAllFiles(word, path string) []FindInfo {
	//findInfos := []FindInfo{}

	//filelist, err := GetFileList(path) //파일 리스트 가져오기
	filelist, err := filepath.Glob(path) //실행 인수 가져오기
	fmt.Println("filelist", filelist)
	if err != nil {
		fmt.Println("파일 경로가 잘못되었습니다. err:", err, "path:", path)
		return findInfos
	}

	ch := make(chan FindInfo) //채널 생성
	cnt := len(filelist)      //파일리스트 길이 체크
	recvCnt := 0

	for _, filename := range filelist {
		go FindWordInFile(word, filename, ch) // 고루틴 실행
	}

	for findInfo := range ch {
		findInfos = append(findInfos, findInfo) //결과 수집
		recvCnt++
		if recvCnt == cnt {
			break
		}
	}
	return findInfos

}

func FindWordInFile(word, filename string, ch chan FindInfo) {
	findInfo := FindInfo{filename, []LineInfo{}}
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("파일을 찾을 수 없습니다. ", filename)
		ch <- findInfo
		return
	}
	defer file.Close()

	lineNo := 1
	scanner := bufio.NewScanner(file) //스캐너를 만든다.
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, word) { //한 줄씩 읽으면 단어 포함 여부 검색
			findInfo.lines = append(findInfo.lines, LineInfo{lineNo, line})

		}
		lineNo++
	}
	ch <- findInfo //채널에 결과 전송
}
