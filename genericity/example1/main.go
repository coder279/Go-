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
	m := make(map[string]int64)
	m["s1"] = 1.0
	sum := SumIntsOrFloats(m)
	log.Printf("sum:%v\n", sum)
}
