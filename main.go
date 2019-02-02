package main

import (
	"context"
	"log"
)

type Input struct {
}

type Output struct {
}

func handler(ctx context.Context, input Input) (*Output, error) {

	log.Printf("handle request")

	return nil, nil
}

func main() {
	lambda.Start(handler)
}
