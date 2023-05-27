package main

import (
	"io"
	"net/url"
)

type Request struct {

	//HTTP 요청 메서드 정보를 가지고 있습니다.
	//"GET", "POST", "PUT", "DELETE"와 같은  메서드입니다.
	Method string

	//HTTP 요청을 보낸 URL 정보를 담고 있습니다.
	//URL 정보를 이용해서 URL에 포함된 데이터를 쿼리해올 수 있습니다.
	URL *url.URL

	//HTTP 프로토콜 버전 정보입니다.
	//HTTP 1.0인지 2.0인지를 알아올 수 있습니다.
	Proto      string //HTTP/1.0
	ProtoMajor int    //1
	ProtoMinor int    //0

	//HTTP 요청 헤더정보입니다.
	//만약 헤더가 다음과 같다면
	/*
		Host: example.com
		accept-encoding: gzip, deflate
		Accept-Language: en-us
		fOO:Bar
		foo:two

		아래와 같이 맵 형태로 변환되어 저장됩니다.
		Header = map[string][]string{
			"Accept-Encoding" : {"gzip, deflate"}
			"Accept-Language" : {"en-us"}
			"Foo":{"Bar","two"}
		}
	*/

	Header Header
	//HTTP 요청의 실제 데이터를 담고 있는 바디 정보입니다.
	//io.Reader 인터페이스를 통해서 데이터를 읽어 올 수 있습니다.
	//io.Reader에 대해서는 A.3절을 참조하세요.

	Body io.ReadCloser
	//그외..
	//http.Request는 이외에도 다양한 정보를 포함하고 있다.
}
