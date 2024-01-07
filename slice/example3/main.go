package main

import "log"

func main() {
	orderLen := 5
	order := make([]uint16, 2*orderLen)
	pollOrder := order[:orderLen:orderLen]
	lockOrder := order[orderLen:][:orderLen:orderLen]
	log.Println(len(pollOrder))
	log.Println(cap(pollOrder))
	log.Println(len(lockOrder))
	log.Println(cap(lockOrder))

}
