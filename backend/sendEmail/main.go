package main

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/yonosoytony/stori-challenge/backend/shared"
)

func HandleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var email struct {
		Email string `json:"email"`
	}
	err := json.Unmarshal([]byte(request.Body), &email)
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 400, Body: err.Error()}, nil
	}

	svc, err := shared.NewDynamoDBClient()
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 500, Body: err.Error()}, nil
	}

	transactions, err := shared.QueryTransactionsByEmail(svc, email.Email)
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 500, Body: err.Error()}, nil
	}

	metrics, err := shared.CalculateMetrics(transactions)
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 500, Body: err.Error()}, nil
	}

	jsonMetrics, err := json.Marshal(metrics)
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 500, Body: err.Error()}, nil
	}

	return events.APIGatewayProxyResponse{StatusCode: 200, Body: string(jsonMetrics)}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
