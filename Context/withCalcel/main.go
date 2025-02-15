package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	cancelCtx, calcelFunc := context.WithCancel(ctx)
	go task(cancelCtx)
	time.Sleep(3 * time.Second)
	calcelFunc()
	time.Sleep(1 * time.Second)
}

func task(ctx context.Context) {
	i := 1
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Gracefully exit")
			fmt.Println(ctx.Err())
			return
		default:
			fmt.Println(i)
			time.Sleep(time.Second * 1)
			i++
		}
	}

}
