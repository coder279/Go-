package main

import (
	"log"
	"time"
)

func main() {
	time.AfterFunc(5*time.Second, func() {
		log.Println("over")
	})
	log.Println("end")
	select {}
}
