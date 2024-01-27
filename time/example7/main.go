package main

import (
	"log"
	"time"
)

func main() {
	log.Println(time.Now())
	<-time.After(1 * time.Second)
	log.Println(time.Now())
}
