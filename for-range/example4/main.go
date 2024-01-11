package main

// 切片在循环的时候已经定义好了
func main() {
	s := []int{1, 2, 3}
	for _, v := range s {
		s = append(s, v)
	}
}
