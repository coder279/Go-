package main

import "time"

func main() {
	ticker := time.Tick(2 * time.Second)
	<-ticker

}
