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
	A   int    `json:"a"`
	B   int    `json:"b"`
	Ops string `json:"ops"`
}
type Output struct {
	Value int `json:"value"`
}

func handler(request Input) (Output, error) {
	fmt.Println("Received: %s\n", request)
	a := (request.A)
	b := (request.B)
	var result int
	switch request.Ops {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		result = a / b
	case "^":
		result = a ^ b
	default:
		return Output{}, errors.New("Unknown Operator")
	}

	return Output{Value: result}, nil
}

func main() {
	lambda.Start(handler)
}
