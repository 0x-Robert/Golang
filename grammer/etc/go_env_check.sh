#!/bin/bash

echo "환경변수 GOOS 확인"
go env GOOS

echo "환경변수 GOARCH를 확인한다."
go env GOARCH

# 환경 변수 OS를 linux로 변경한다.
#export GOOS="linux"

# GO의 환경 변수 GOARCH를 386으로 변경한다.
# export GOARCH="386"

# GOOS와 GOARCH에 지정된 환경을 타겟으로 빌드한다.
# go build
