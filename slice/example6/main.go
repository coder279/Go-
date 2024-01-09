package main

import "log"

func main() {
	simple := []int{1, 2, 3, 4, 5}
	//切片后append会覆盖原来的数值,封印切片如果切片增加数据后则产生新的切片
	simpleA := simple[1:3:3]
	log.Printf("simpleA:%p simple:%p", simpleA, simple)
	simpleA = append(simpleA, 10)
	log.Println(simple)
}
