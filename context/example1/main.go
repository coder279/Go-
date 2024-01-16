package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func HandleRequest(ctx context.Context) {
	go HandleWriteMysql(ctx)
	go HandleWriteRedis(ctx)
	for {
		select {
		case <-ctx.Done():
			fmt.Println("HandelRequest Done.")
			return
		default:
			fmt.Println("HandelRequest running")
			time.Sleep(2 * time.Second)
		}
	}

}
func HandleWriteMysql(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("write mysql done")
			return
		default:
			fmt.Println("write mysql writing")
			time.Sleep(2 * time.Second)
		}
	}
}
func HandleWriteRedis(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("write redis done")
			return
		default:
			fmt.Println("write redis writing")
			time.Sleep(2 * time.Second)
		}
	}
}
func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go HandleRequest(ctx)
	time.Sleep(time.Second * 5)
	cancel()
	time.Sleep(time.Second * 5)
	log.Println("end...")
}
