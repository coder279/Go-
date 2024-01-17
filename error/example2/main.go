package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	err := errors.New("not found")
	err1 := fmt.Errorf("some context:%v", err)
	err2 := fmt.Errorf("some context:%v", err)
	log.Println(err1 == err2)
}
