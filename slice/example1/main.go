package main

import "log"

func sliceRice(s []int) {
	s = append(s, 0)
	for i := range s {
		s[i]++
	}
}
func main() {
	s1 := []int{1, 2}
	s2 := s1
	s2 = append(s2, 3)
	sliceRice(s1)
	sliceRice(s2)
	log.Println(s1, s2)
}
