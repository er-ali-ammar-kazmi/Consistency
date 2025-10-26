package problems

import (
	"fmt"
	"math/rand/v2"
	"strconv"
	"time"
)

func Start(duration float64) {
	// Example to run producer-consumer
	start := time.Now()
	done := make(chan struct{})
	fn := func() any {
		return strconv.QuoteRuneToASCII(rune(rand.IntN(120)))
	}

	stream := Producer(done, fn)
	Consumer(done, stream)
	for {
		since := time.Since(start)
		if since.Seconds() > duration {
			close(done)
			time.Sleep(time.Second)
			return
		}
	}
}

func Producer[T any](done <-chan struct{}, fn func() T) <-chan T {
	stream := make(chan T)
	go func() {
		defer close(stream)
		for {
			time.Sleep(time.Second * 1 / 2)
			select {
			case <-done:
				fmt.Println("Producer Closing")
				return
			case stream <- fn():
			}
		}
	}()

	return stream
}

func Consumer[T any](done <-chan struct{}, stream <-chan T) {
	go func() {
		for {
			select {
			case <-done:
				fmt.Println("Consumer Closing")
				return
			case num := <-stream:
				fmt.Println(num)
			}
		}
	}()
}
