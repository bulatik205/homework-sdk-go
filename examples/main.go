package main

import (
	"context"
	"fmt"
	"log"

	homework "github.com/bulatik205/homework-sdk-go"
)

func main() {
	client := homework.New(
		homework.WithBaseURL("..."),
	)

	if err := client.Ping(context.Background()); err != nil {
		log.Fatal(err)
	}
	fmt.Println("pong")
}
