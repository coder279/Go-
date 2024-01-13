package main

import (
	"log"
	"runtime"
)

func main() {
	log.Println(runtime.GOMAXPROCS(80)) //设置处理器，每个P都维护一个G
}
