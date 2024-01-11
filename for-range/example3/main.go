package main

import (
	"fmt"
	"time"
)

// 1s后输出结束后，陷入阻塞
func main() {
	t := time.NewTimer(time.Second)
	for _ = range t.C {
		fmt.Println("hi")
	}
}
