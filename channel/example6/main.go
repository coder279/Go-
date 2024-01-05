package main

import (
	"log"
	"time"
)

func main() {
	ch := make(chan int, 2)
	go func() {
		for {
			ch <- 1
			time.Sleep(2 * time.Second)
		}
	}()
	for e := range ch {
		log.Printf("Get element from chan:%d\n", e)
	}
}
