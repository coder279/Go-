package main

import (
	"log"
	"time"
)

func main() {
	go func() {
		timer := time.NewTimer(1 * time.Second)
		select {
		case <-timer.C:
			log.Println("event")
			timer.Reset(1 * time.Second)
		}
	}()
	select {}
}
