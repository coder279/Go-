package main

import (
	"errors"
	"log"
)

func main() {
	err1 := errors.New("not found")
	err2 := errors.Unwrap(err1)
	log.Println(err1, err2)
}
