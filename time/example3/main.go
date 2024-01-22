package main

import (
	"log"
	"time"
)

func AfterDemo() {
	log.Println(time.Now())
	<-time.After(1 * time.Second)
	log.Println(time.Now())
}

func AfterFuncDemo() {
	log.Println("AfterFuncDemo start: ", time.Now())
	time.AfterFunc(1*time.Second, func() {
		log.Println("AfterFuncDemo end: ", time.Now())
	})
	time.Sleep(2 * time.Hour)
}

func main() {
	AfterDemo()
	AfterFuncDemo()
}
