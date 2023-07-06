package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	//클라이언트가 서버와 연결을 시도한다.
   conn, err := net.Dial("tcp", ":8000")
   if nil != err {
      log.Println(err)
   }

   //고루틴을 생성해서 서버가 값을 던질때까지 기다렸다가 던지면 값을 출력한다.
   go func() {
      data := make([]byte, 4096)

      for {
         n, err := conn.Read(data)
         if err != nil {
            log.Println(err)
            return
         }

         log.Println("Server send : " + string(data[:n]))
         time.Sleep(time.Duration(3) * time.Second)
      }
   }()

   //사용자의 입력이 들어올때까지 blocking했다가 입력을 마치면 서버로 전송한다.
   for {
      var s string
      fmt.Scanln(&s)
      conn.Write([]byte(s))
      time.Sleep(time.Duration(3) * time.Second)
   }
}