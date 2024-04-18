package main

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// Your Lambda function logic here
	return events.APIGatewayProxyResponse{
		Body:       "Hello from Lambda!",
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}
