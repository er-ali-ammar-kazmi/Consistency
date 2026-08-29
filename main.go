package main

import (
	"fmt"
	"os"

	env "github.com/joho/godotenv"
)

func main() {
	env.Load()

	fmt.Println("Hello " + os.Getenv("SECRET"))
}
