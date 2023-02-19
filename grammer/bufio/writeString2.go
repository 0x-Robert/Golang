// os.Create() 를 이용하여 파일 생성 후 fmt.Fprintf() 를 사용하여 쓰는 방식
// os.Create() 를 이용하여 파일 생성 후 f.WriteString(), f.Write() 를 사용하여 쓰는 방식
// os.Create() 를 이용하여 파일 생성 후 bufio.NewWriter() 를 사용하여 쓰는 방식
// ioutil.WriteFile() 를 사용하여 쓰는 방식
// os.Create() 를 이용하여 파일 생성 후 io.WriteString() 를 사용하여 쓰는 방식

package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)
func checkError(err error) {
  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
}

func main() {
  b := []byte("Data to write\n")
  // 1.
  f1, err := os.Create("/tmp/f1.txt")
  checkError(err)
  defer f1.Close()
  fmt.Fprintf(f1, string(b))
  
  // 2.
  f2, err := os.Create("/tmp/f2.txt")
  checkError(err)
  defer f2.Close()
  n, err := f2.WriteString(string(b))
  n, err := f2.Write(b) <br>  fmt.Printf("wrote %d bytes\n", n)

  // 3.
  f3, err := os.Create("/tmp/f3.txt")
  checkError(err)
  w := bufio.NewWriter(f3)
  n, err = w.WriteString(string(b))
  fmt.Printf("wrote %d bytes\n", n)
  w.Flush()
  
  // 4.
  f4 := "/tmp/f4.txt"
  err = ioutil.WriteFile(f4, b, 0644)
  checkError(err)
  // 5.
  f5, err := os.Create("/tmp/f5.txt")
  checkError(err)
  n, err = io.WriteString(f5, string(b))
  checkError(err)
  fmt.Printf("wrote %d bytes\n", n)
}