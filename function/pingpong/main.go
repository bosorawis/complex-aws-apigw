package main

import (
	"errors"
	"fmt"
	"strconv"

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
	A   string `json:"a"`
	B   string `json:"b"`
	Ops string `json:"ops"`
}
type Output struct {
	Value string `json:"value"`
}

func handler(request Input) (Output, error) {
	fmt.Println("Received: %s\n", request)
	a, _ := strconv.Atoi(request.A)
	b, _ := strconv.Atoi(request.B)
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

	strAns := strconv.Itoa(result)
	return Output{Value: strAns}, nil
}

func main() {
	lambda.Start(handler)
}
