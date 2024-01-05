package main

import "log"

func main() {
	//使用channel实现互斥锁,有一个缓存的可以形成互斥锁
	ch := make(chan struct{}, 1)
	//创建一个map,map并发情况会报错
	bm := make(map[string]int)
	go func() {
		for {
			ch <- struct{}{}
			bm["cool"]++
			<-ch
			log.Printf("result:%d", bm["cool"])
		}
	}()
	select {}

}
