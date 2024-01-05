package main

import (
	"log"
	"time"
)

func addNumberChan(data chan struct{}) {
	for {
		data <- struct{}{}
		time.Sleep(1 * time.Second)
	}
}
func main() {
	//select顺序是随机的
	var ch1 = make(chan struct{}, 1)
	var ch2 = make(chan struct{}, 1)
	go addNumberChan(ch1)
	go addNumberChan(ch2)
	for {
		select {
		case <-ch1:
			log.Println("chan1 value ")
		case <-ch2:
			log.Println("chan2 value ")
		default:
			log.Println("nothing")
			time.Sleep(1 * time.Second)
		}
	}

}
