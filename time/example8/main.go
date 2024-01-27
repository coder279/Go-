package main

import (
	"log"
	"time"
)

func main() {
	log.Println("start")
	time.AfterFunc(1*time.Second, func() {
		log.Println("end..")
	})
	time.Sleep(2 * time.Second)
}
