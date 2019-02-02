package main

import (
	"net/http"

	"github.com/aws/aws-lambda-go/lambda"
)

type Response struct {
	IsBase64Encoded bool              `json:"isBase64Encoded"`
	StatusCode      int               `json:"statusCode"`
	Header          map[string]string `json:"headers"`
	Body            interface{}       `json:"body"`
}

func hello() (Response, error) {
	return Response{
		IsBase64Encoded: false,
		StatusCode:      http.StatusOK,
		Header:          map[string]string{},
		Body:            "Hello World",
	}, nil
}

func main() {
	lambda.Start(hello)
}
