package main

import (
	"fmt"
	"time"
)

func ping(channel chan string) {
	for i := 0; ; i++ {
		channel <- "ping"
	}
}

func pong(channel chan string) {
	for i := 0; ; i++ {
		channel <- "pong"
	}
}

func print(channel chan string) {
	for {
		msg := <-channel
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}

func main() {
	channel := make(chan string)

	go ping(channel)
	go pong(channel)
	go print(channel)

	var enter string
	fmt.Scanln(&enter)
}