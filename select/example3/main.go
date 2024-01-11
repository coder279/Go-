package main

import "log"

func main() {
	c := make(chan int, 10)
	c <- 1
	close(c)
	select {
	case d := <-c:
		log.Println(d)
	}
}
