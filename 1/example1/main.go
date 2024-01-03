package main

func main() {
	//读关闭的channel
	ch := make(chan int)
	close(ch)
	//读关闭的channel不会发生panic
	<-ch
	//写关闭的channel会发生panic
	ch <- 1
	//读，写为nil的channel也不会发生panic只是会发生阻塞
}
