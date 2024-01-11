package main

import "log"

func main() {
	c := make(chan int, 10)
	c <- 1
	close(c)
	select {
	case d, ok := <-c:
		if !ok {
			log.Println("no data received")
			break
		}
		log.Println(d, ok)
	}

}
