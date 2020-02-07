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

type Input struct {
	a   int    `json:"a"`
	b   int    `json:"b"`
	ops string `json:"ops"`
}
type Output struct {
	value int `json:"value"`
}

func handler(request Input) (Output, error) {
	fmt.Printf("Received: %s\n", request.Message)
	switch request.ops {
	case "+":
		return Output{value: a + b}, nil
	case "-":
		return Output{value: a - b}, nil
	case "*":
		return Output{value: a * b}, nil
	case "/":
		return Output{value: a / b}, nil
	case "^":
		return Output{value: a ^ b}, nil
	default:
		return nil, errors.New("Unknown Operator %w")
	}
}

func main() {
	lambda.Start(handler)
}
