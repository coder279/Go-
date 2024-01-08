package main

import "log"

func main() {
	s := *new([]int)
	log.Println(s)
	s = append(s, []int{5, 6}...)
	log.Println(s)
}
