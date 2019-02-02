package main

import (
	"context"
	"log"
	"strings"

	"github.com/aws/aws-lambda-go/lambda"
)

func handler(ctx context.Context, input string) (string, error) {
	log.Printf("handle requtest for %s", input)
	return strings.ToUpper(input), nil
}

func main() {
	lambda.Start(handler)
}
