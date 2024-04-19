package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type Transaction struct {
	ID        string  `json:"id"`
	Email     string  `json:"email"`
	Month     int     `json:"month"`
	Amount    float64 `json:"amount"`
	Timestamp string  `json:"timestamp"`
}

type RequestBody struct {
	Email  string  `json:"email"`
	Amount float64 `json:"amount"`
	Month  int     `json:"month"`
}

func HandleRequest(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var requestBody RequestBody
	err := json.Unmarshal([]byte(event.Body), &requestBody)
	if err != nil {
		errorMsg := fmt.Sprintf("Error unmarshalling request body: %s", err.Error())
		log.Println(errorMsg)
		return events.APIGatewayProxyResponse{StatusCode: 400, Body: errorMsg}, nil
	}

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	svc := dynamodb.New(sess)

	// Generate a unique ID combining timestamp and email
	timestamp := time.Now().Format(time.RFC3339)
	id := fmt.Sprintf("%s-%s", timestamp, requestBody.Email)

	transaction := Transaction{
		ID:        id,
		Email:     requestBody.Email,
		Month:     requestBody.Month,
		Amount:    requestBody.Amount,
		Timestamp: timestamp,
	}

	av, err := dynamodbattribute.MarshalMap(transaction)
	if err != nil {
		errorMsg := fmt.Sprintf("Error marshalling transaction: %s", err.Error())
		log.Println(errorMsg)
		return events.APIGatewayProxyResponse{StatusCode: 500, Body: errorMsg}, nil
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String("stori-challenge-transactions"),
	}

	_, err = svc.PutItem(input)
	if err != nil {
		errorMsg := fmt.Sprintf("Error putting item into DynamoDB: %s", err.Error())
		log.Println(errorMsg)
		return events.APIGatewayProxyResponse{StatusCode: 500, Body: errorMsg}, nil
	}

	// Return a successful response
	responseBody, err := json.Marshal(map[string]string{
		"message": "Transaction successfully added",
	})
	if err != nil {
		errorMsg := fmt.Sprintf("Error marshalling response: %s", err.Error())
		log.Println(errorMsg)
		return events.APIGatewayProxyResponse{StatusCode: 500, Body: errorMsg}, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(responseBody),
	}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
