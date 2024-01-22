package main

import (
	"log"
	"time"
)

func WaitChannel(conn <-chan string) bool {
	timer := time.NewTimer(1 * time.Second)
	select {
	case <-conn:
		timer.Stop()
	case <-timer.C:
		log.Println("WaitChannel timeout!!")
	}
	return false
}
func main() {
	ch := make(chan string, 1024)
	WaitChannel(ch)
}
