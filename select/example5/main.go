package main

import (
	"log"
	"time"
)

func main() {
	c := make(chan string, 1)
	c <- "nice"
	var recv string
	send := "hello"
	for {
		select {
		case recv = <-c:
			log.Println(recv)
		case c <- send:
			log.Println(send)
		}
		time.Sleep(1 * time.Second)
	}
}
