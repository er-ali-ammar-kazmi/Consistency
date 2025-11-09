package program

import (
	"context"
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

func Start(d int) {
	// Example to run producer-consumer

	var wg sync.WaitGroup

	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Second*time.Duration(d)))
	defer cancel()

	fn := func() any {
		time.Sleep(time.Second / 2)
		return string(rune(rand.IntN(118)))
	}

	stream := Producer(ctx, &wg, fn)
	Consumer(&wg, stream)

	wg.Wait()
}

func Producer[T any](ctx context.Context, wg *sync.WaitGroup, fn func() T) <-chan T {
	stream := make(chan T)
	(*wg).Go(func() {
		defer close(stream)
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Producer Closing : ", ctx.Err().Error())
				return
			case stream <- fn():
			}
		}
	})

	return stream
}

func Consumer[T any](wg *sync.WaitGroup, stream <-chan T) {
	(*wg).Go(func() {
		for num := range stream {
			fmt.Print(num, " ")
		}
		fmt.Println("Consumer Closing!")
	})
}
