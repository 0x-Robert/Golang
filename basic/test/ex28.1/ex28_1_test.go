package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSquare1(t *testing.T) {
	assert := assert.New(t)                               //테스트 객체생성
	assert.Equal(81, square(9), "square(9) should be 81") //테스트 함수 호출
}

func TestSquare2(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(9, square(3), "square(3) should be 9")

}
