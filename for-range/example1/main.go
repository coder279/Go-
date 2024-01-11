package main

import (
	"fmt"
	"sync"
)

func PrintSlice() {
	s := []int{1, 2, 3}
	var wg sync.WaitGroup
	wg.Add(len(s))
	for _, v := range s {
		v := v
		go func() {
			fmt.Println(v)
			wg.Done()
		}()
	}
	wg.Wait()
}
func main() {
	PrintSlice()
}
