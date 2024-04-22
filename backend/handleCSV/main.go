package main

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/yonosoytony/stori-challenge/backend/shared"
)

func HandleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var payload struct {
		Transactions []shared.Transaction `json:"transactions"`
	}
	err := json.Unmarshal([]byte(request.Body), &payload)
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 400, Body: err.Error()}, nil
	}

	svc, err := shared.NewDynamoDBClient()
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 500, Body: err.Error()}, nil
	}

	err = shared.PutTransactions(svc, payload.Transactions)
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 500, Body: err.Error()}, nil
	}

	return events.APIGatewayProxyResponse{StatusCode: 200, Body: "Transactions created successfully"}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
