// https://juejin.im/post/6844904070667321357

package main

import (
	"fmt"
	"time"

	"context"
)

func main() {
	messages := make(chan int, 10)

	// producer
	for i := 0; i < 10; i++ {
		messages <- i
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	go func(ctx context.Context) {
		ticker := time.NewTicker(1 * time.Second)
		for _ = range ticker.C {
			select {
			case <-ctx.Done():
				fmt.Println("child process interrupt....")
				return
			default:
				fmt.Printf("send messages: %d\n", <-messages)
			}
		}
	}(ctx)

	defer close(messages)
	defer cancel()

	select {
	case <-ctx.Done():
		time.Sleep(1 * time.Second)
		fmt.Println("main process exit")
	}
}
