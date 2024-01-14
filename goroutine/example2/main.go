package main

import (
	"log"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(1)
	go func() {
		for {
		}
	}()
	time.Sleep(1 * time.Second)
	log.Println("DONE")

}
