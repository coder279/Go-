package main

import "log"

func sliceCap() {
	var arr = [10]int{}
	var slice = arr[5:6]
	//len取的是当前slice数据长度
	log.Printf("len slice = %d\n", len(slice))
	//cap取得是array截取后的数据长度
	log.Printf("cap slice = %d\n", cap(slice))
}
func main() {
	sliceCap()
}
