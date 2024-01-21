package main

import (
	"log"
	"sync"
)

type Set[T comparable] map[T]struct{}

type ThreadSafeSet[T comparable] struct {
	l sync.Mutex
	m map[T]struct{}
}

func MakeSet[T comparable]() Set[T] {
	return make(Set[T])
}

// Add 添加元素
func (s Set[T]) Add(v T) {
	s[v] = struct{}{}
}

// Delete 删除元素
func (s Set[T]) Delete(v T) {
	delete(s, v)
}

// Contains 包含元素
func (s Set[T]) Contains(v T) bool {
	_, ok := s[v]
	return ok
}

func (s Set[T]) Len() int {
	return len(s)
}

func (s Set[T]) Iterate(f func(T)) {
	for v := range s {
		f(v)
	}
}
func main() {
	set := MakeSet[int]()
	set.Add(1)
	if set.Contains(2) {
		log.Println("insert success")
	}

}
