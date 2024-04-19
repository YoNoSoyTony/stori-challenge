package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type Transaction struct {
	ID     string  `json:"id"`
	Email  string  `json:"email"`
	Month  int     `json:"month"`
	Amount float64 `json:"amount"`
}

type RequestBody struct {
	Email string `json:"email"`
}

func HandleRequest(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var requestBody RequestBody
	marshall_err := json.Unmarshal([]byte(event.Body), &requestBody)
	if marshall_err != nil {
		errorMsg := fmt.Sprintf("Error unmarshalling request body: %s", marshall_err.Error())
		log.Println(errorMsg)
		return events.APIGatewayProxyResponse{StatusCode: 400, Body: errorMsg}, nil
	}

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	svc := dynamodb.New(sess)

	var result *dynamodb.QueryOutput
	var err error // Declare err here

	if requestBody.Email != "" {
		// Query by email
		queryInput := &dynamodb.QueryInput{
			TableName:              aws.String("stori-challenge-transactions"),
			IndexName:              aws.String("EmailIndex"), // Ensure this matches your GSI name
			KeyConditionExpression: aws.String("Email = :email"),
			ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
				":email": {
					S: aws.String(requestBody.Email),
				},
			},
		}
		result, err = svc.Query(queryInput) // Assign new value to err
	} else {
		// Scan all transactions
		scanInput := &dynamodb.ScanInput{
			TableName: aws.String("stori-challenge-transactions"),
		}
		var scanResult *dynamodb.ScanOutput
		scanResult, err = svc.Scan(scanInput) // Assign new value to err
		if err != nil {
			errorMsg := fmt.Sprintf("Error scanning transactions: %s", err.Error())
			log.Println(errorMsg)
			return events.APIGatewayProxyResponse{StatusCode: 500, Body: errorMsg}, nil
		}
		result = &dynamodb.QueryOutput{Items: scanResult.Items}
	}

	if err != nil {
		errorMsg := fmt.Sprintf("Error querying transactions: %s", err.Error())
		log.Println(errorMsg)
		return events.APIGatewayProxyResponse{StatusCode: 500, Body: errorMsg}, nil
	}

	// Unmarshal the result
	var transactions []Transaction
	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &transactions) // Assign new value to err
	if err != nil {
		errorMsg := fmt.Sprintf("Error unmarshalling transactions: %s", err.Error())
		log.Println(errorMsg)
		return events.APIGatewayProxyResponse{StatusCode: 500, Body: errorMsg}, nil
	}

	// Return the transactions
	responseBody, err := json.Marshal(transactions) // Assign new value to err
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
