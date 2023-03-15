package main

import (
	"io"
	"log"
	"net"
)

func main() {
	//소켓을 열어준다.
   l, err := net.Listen("tcp", ":8000")

   if nil != err {
      log.Println(err);
   }

   //메인 프로세스가 종료되면 소켓이 종료되게 설정한다.
   defer l.Close()

   for {
      conn, err := l.Accept()
      if nil != err {
         log.Println(err);
         continue
      }
	  //연결은 메인프로세스종료시 종료되도록 해준다.
      defer conn.Close()

	  //소켓에 연결하고 그 연결을 우리가 만든 
	  //connHandler에 인자로 주고 고루틴으로 돌린다.
      go ConnHandler(conn)
   }
}

func ConnHandler(conn net.Conn) {
   recvBuf := make([]byte, 4096)
   for {
	//client가 값을 줄때까지 blocking되어 대기하다가 값을 주면 읽어들인다.
      n, err := conn.Read(recvBuf)
      if nil != err {
         if io.EOF == err {
            log.Println(err);
            return
         }
         log.Println(err);
         return
      }
	  //client가 던진 값을 다시 client에게 던진다.
      if 0 < n {
         data := recvBuf[:n]
         log.Println(string(data))
         _, err = conn.Write(data[:n])
         if err != nil {
            log.Println(err)
            return
         }
      }
   }
}