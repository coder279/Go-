package main

import (
	"log"
	"time"
)

func main() {
	timer := time.NewTimer(5 * time.Second)
	<-timer.C
	log.Println(timer.Stop())
	log.Println("success")
}
