package main

import "log"

func main() {
	c := make(chan string)
	close(c)
	selectAssign(c)

}

func selectAssign(c chan string) {
	select {
	case <-c:
		log.Println("0")
	case d := <-c:
		log.Println("1", d)
	case d, ok := <-c:
		if !ok {
			log.Println("close")
			break
		}
		log.Println("2", d)

	}
}
