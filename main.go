package main

import (
	"consistency/program"
	"fmt"

	env "github.com/joho/godotenv"
)

func main() {
	env.Load()

	fmt.Println("Hello World")
	program.DaisyChain()
}
