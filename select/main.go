package main

import "log"

func main() {
	c := make(chan int)
	select {
	case <-c:
		log.Println("read")

	case c <- 1:
		log.Println("write")
	}

}
