#!/bin/bash

#localhost.key에 비밀키 저장
#localhost.csr에 인증파일 저장 > 이 인증파일을 인증기관에 제출해서 인증서인 .crt 파일을 생성함 
openssl req -new -newkey rsa:2048 -nodes -keyout localhost.key -out localhost.csr 

