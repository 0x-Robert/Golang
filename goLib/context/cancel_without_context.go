package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	exit := make(chan os.Signal)
	signal.Notify(exit, syscall.SIGINT, syscall.SIGTERM)

	quit := make(chan struct{})

	go func() {
		fmt.Println("Signal:", <-exit)
		quit <- struct{}{}

	}()

	start := time.Now()
	result, err := longFuncWithoutCtx(quit)
	fmt.Printf("duration:%v result:%s\n", time.Since(start), result)
	if err != nil {
		log.Fatal(err)
	}

}

func longFuncWithoutCtx(quit <-chan struct{}) (string, error) {
	done := make(chan string)

	go func() {
		done <- longFunc()
	}()

	select {
	case <-quit:
		return "Fail", errors.New("Force quit")
	case result := <-done:
		return result, nil
	}
}
func longFunc() string {
	<-time.After(time.Second * 3)
	return "Success"
}
