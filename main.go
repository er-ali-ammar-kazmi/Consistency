package main

import (
	"practise/program"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	program.Start(6)
}
