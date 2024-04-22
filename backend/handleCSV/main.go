package main

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/yonosoytony/stori-challenge/backend/shared"
)

func HandleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var transaction shared.Transaction
	err := json.Unmarshal([]byte(request.Body), &transaction)
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 400, Body: err.Error()}, nil
	}

	transaction.GenerateTransactionID()

	svc, err := shared.NewDynamoDBClient()
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 500, Body: err.Error()}, nil
	}

	err = shared.PutTransaction(svc, transaction)
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 500, Body: err.Error()}, nil
	}

	return events.APIGatewayProxyResponse{StatusCode: 200, Body: "Transaction created successfully"}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
