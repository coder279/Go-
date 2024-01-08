package main

import "log"

func main() {
	simple := []int{1, 2, 3, 4, 5}
	//切片后append会覆盖原来的数值
	simpleA := simple[1:3]
	simpleA = append(simpleA, 10)
	log.Println(simple)
}
