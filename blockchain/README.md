<!-- 먼저 블록체인을 공부하기 위해서는 다음과 같은 과정이 필요합니다.
나만의 블록체인 만들어 보기
해싱 워크가 어떻게 블록체인의 무결성을 유지시키는지에 대한 이해
어떻게 새로운 블락들이 생기는지 알기
여러 노드에서 동시에 블록이 생성될때 이를 어떻게 해결하는지에 대한 이해
웹브라우저에서 블록체인 확인
새로운 블록 작성
블록체인 시작
Go lang을 사용하여 블록체인을 만들어 봅시다.

환경 설정
먼저 Go lang을 설치하고 configure를 설정 해야하는데 다음 패키지들을 설치해야합니다.

Spew 설치
go get github.com/davecgh/go-spew/spew

Spew는 콘솔에서 Go data구조와 각각의 부분을 완벽한 포맷으로 볼수 있게 해주는 패키지 입니다. 디버깅을 할때 매우 유용합니다.

Gorilla/mux설치
go get github.com/gorilla/mux
Gorilla/mux 는 웹 핸들러를 작성하기 위해 많이 사용하는 패키지입니다. 

Gotdotenv
go get github.com/joho/godotenv
Gotdotenv는 .env파일(root 디렉토리에서 http port같이 하드코드를 할 수 있는 파일)로 부터 읽을수 있게 해주는 패키지 입니다.

소스가 있는 root 디렉토리에 .env 파일을 생성하고 다음과 같이 http 포트 작성하면 됩니다.

ADDR=8080

main.go 파일을 생성하시고 코딩을 시작해봅시다. -->