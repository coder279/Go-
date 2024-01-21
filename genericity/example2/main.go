package main

import "log"

type Vector[T any] []T //声明一个切片来存储元素，元素类型是any

func (v *Vector[T]) push(x T) {
	*v = append(*v, x)
}
func ExampleVectorPush() {
	var v Vector[string]
	v.push("Hello World")
	log.Println(v)
}
func main() {
	ExampleVectorPush()
}
