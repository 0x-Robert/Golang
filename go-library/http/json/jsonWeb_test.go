package main

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestJsonHandler(t *testing.T) {
	assert := assert.New(t)

	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/student", nil) //student 경로 테스트

	mux := MakeWebHandler()
	mux.ServeHTTP(res, req)

	assert.Equal(http.StatusOK, res.Code)
	student := new(Student)
	err := json.NewDecoder(res.Body).Decode(student) //결과 반환
	assert.Nil(err)                                  //결과 확인
	assert.Equal("aaa", student.Name)
	assert.Equal(16, student.Age)
	assert.Equal(87, student.Score)
}
