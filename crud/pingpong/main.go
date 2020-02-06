package main

import (
	"errors"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
)

var (
	// DefaultHTTPGetAddress Default Address
	DefaultHTTPGetAddress = "https://checkip.amazonaws.com"

	// ErrNoIP No IP found in response
	ErrNoIP = errors.New("No IP in HTTP response")

	// ErrNon200Response non 200 status code in response
	ErrNon200Response = errors.New("Non 200 Response found")
)

type MSG struct {
	Message string `json:"message"`
}


func handler(request MSG) (MSG, error) {
	fmt.Printf("Received: %s\n", request.Message)
	return MSG{Message: fmt.Sprintf("pong - %s", request.Message)}, nil
}

func main() {
	lambda.Start(handler)
}
