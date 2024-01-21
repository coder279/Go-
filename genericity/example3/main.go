package main

import "log"

// 声明一个string的近似类型
type Mystring interface {
	~string
}

type Integer interface {
	~int
}

func sayHi[T Mystring](name T) {
	log.Printf("Hi %s", name)
}
func count[T Integer](number T) {
	log.Printf("number:%d", number)
}
func main() {
	type s string
	var n s = "Leo"
	sayHi(n)

	var number1 int
	number1 = 10
	count(number1)

	//var number2 int8
	//number2 = 10
	//count(number2)
}
