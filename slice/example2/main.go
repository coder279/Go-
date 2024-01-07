package main

import "log"

func sliceExpend() {
	var slice []int
	s1 := append(slice, 1, 2, 3, 4)
	log.Printf("len:%d cap:%d point:%p", len(s1), cap(s1), s1)
	s2 := append(s1, 4)
	log.Printf("len:%d cap:%d point:%p", len(s2), cap(s2), s2)
	log.Println(&s1[0] == &s2[0])
}
func main() {
	sliceExpend()
}
