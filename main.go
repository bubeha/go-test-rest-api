package main

import "github.com/bubeha/go-test-rest-api/app"

func main() {
	application := app.New()

	application.Init()

	application.Start()
}
