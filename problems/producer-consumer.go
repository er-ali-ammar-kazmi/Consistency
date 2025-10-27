package problems

import (
	"context"
	"fmt"
	"math/rand/v2"
	"strconv"
	"sync"
	"time"
)

func Start(d int) {
	// Example to run producer-consumer

	var wg sync.WaitGroup

	c := context.Background()

	ctx, _ := context.WithDeadline(c, time.Now().Add(time.Second*time.Duration(d)))

	fn := func() any {
		time.Sleep(time.Second / 2)
		return strconv.QuoteRuneToASCII(rune(rand.IntN(118)))
	}

	stream := Producer(ctx, &wg, fn)
	Consumer(ctx, &wg, stream)
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

func Consumer[T any](ctx context.Context, wg *sync.WaitGroup, stream <-chan T) {
	(*wg).Go(func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println()
				fmt.Println("Consumer Closing : ", ctx.Err().Error())
				return
			case num := <-stream:
				fmt.Print(num)
			}
		}
	})
}
