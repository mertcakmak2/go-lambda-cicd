package main

import (
	"context"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(ctx context.Context) (string, error) {
	return "Merhaba, bu bir Lambda fonksiyonundan dönen string değerdir!", nil
}

func main() {
	lambda.Start(handler)
}
