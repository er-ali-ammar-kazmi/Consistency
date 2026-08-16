package program

import (
	"fmt"
	"math/rand"
)

func DaisyChain() {
	count := rand.Intn(100)
	fmt.Println("Started Chain Generation for Count: ", count)
	listener := make(chan int)
	left := listener
	right := listener

	for i := 0; i < count; i++ {
		right = make(chan int)
		go func(left, right chan int) {
			left <- 1 + <-right
		}(left, right)

		left = right
	}

	right <- 1

	fmt.Println("Listened Count: ", <-listener)
}
