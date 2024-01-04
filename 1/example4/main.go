package main

import "log"

func main() {
	var ch = make(chan int, 1)
	ch <- 1
	v, ok := <-ch
	log.Printf("channel的状态:%v,channel的值:%v", ok, v)
	close(ch)
	v, ok = <-ch
	//对于一个已经关闭的channel读取的时候会返回0值，第二个状态返回false
	log.Printf("channel的状态:%v,channel的值:%v", ok, v)

}
