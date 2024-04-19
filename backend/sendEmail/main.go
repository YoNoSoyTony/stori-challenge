package main

import (
	"context"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type Transaction struct {
	Email  string  `json:"email"`
	Month  int     `json:"month"`
	Amount float64 `json:"amount"`
}

func HandleRequest(ctx context.Context) error {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	svc := dynamodb.New(sess)

	transaction := Transaction{
		Email:  "yonosoytony@duck.com",
		Month:  7,
		Amount: -17,
	}

	av, err := dynamodbattribute.MarshalMap(transaction)
	if err != nil {
		return err
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String("transactions"),
	}

	_, err = svc.PutItem(input)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	lambda.Start(HandleRequest)
}
