package main

import "log"

func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}
func main() {
	m := make(map[string]float64)
	m["s1"] = 1.0
	//sum := SumIntsOrFloats[string, float64](m) //显示调用
	sum := SumIntsOrFloats(m) //隐式调用
	log.Printf("sum:%v\n", sum)
}
