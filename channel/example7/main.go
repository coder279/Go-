package main

import (
	"log"
	"time"
)

func main() {
	ch := make(chan int, 1)
	go func() {
		for {
			ch <- 1
			time.Sleep(20 * time.Second)
		}
	}()
	select {
	case <-ch:
		chanRange(ch)
	case <-time.After(10 * time.Second):
		log.Println("time out")
	}
	log.Println("end...")
}

func chanRange(changeName chan int) {
	for e := range changeName {
		log.Printf("Get element from chan:%d\n", e)
	}
}
