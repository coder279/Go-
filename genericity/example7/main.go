package main

import "sort"

type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64 |
		~string
}

type orderedSlice[T Ordered] []T

func OrderedSlice[T Ordered](s []T) orderedSlice[T] {
	sort.Sort(orderedSlice[T](s))
	return s
}

func (s orderedSlice[T]) Len() int {
	return len(s)
}

func (s orderedSlice[T]) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s orderedSlice[T]) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func main() {
	s := []int{3, 5, 2}
	OrderedSlice(s)
}
