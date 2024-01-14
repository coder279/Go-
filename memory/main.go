package main

import "fmt"

func Fibonac1() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}
func main() {
	f := Fibonac1()
	for i := 0; i < 10; i++ {
		fmt.Printf("Fibnocaci:%d\n", f())
	}
}
