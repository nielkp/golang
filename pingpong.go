package main

import (
	"fmt"
	"time"
)

func pingPong(ch chan<- string) {
	for i := 0; i < 5; i++ {
		ch <- "ping"
		time.Sleep(time.Second)
	}
	close(ch)
}
func pongPing(ch <-chan string) {
	for msg := range ch {
		fmt.Println(msg)
		time.Sleep(time.Second)
		fmt.Println("pong")
	}
}
func main() {
	ch := make(chan string)
	go pingPong(ch)
	pongPing(ch)
}
