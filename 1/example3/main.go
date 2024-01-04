package main

import "log"

func main() {
	//对值为nil的channel读写都会阻塞
	var ch chan int
	ch <- 1
	log.Fatalln("进行下去。。。")
}
