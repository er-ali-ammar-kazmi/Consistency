package main

import (
	app "practise/application"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	app.StartRestCalculatorServer()
}
