package shared

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

func NewDynamoDBClient() (*dynamodb.DynamoDB, error) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"),
	})
	if err != nil {
		return nil, err
	}
	return dynamodb.New(sess), nil
}

func PutTransaction(svc *dynamodb.DynamoDB, transaction Transaction) error {
	item, err := dynamodbattribute.MarshalMap(transaction)
	if err != nil {
		return err
	}

	input := &dynamodb.PutItemInput{
		Item:      item,
		TableName: aws.String("stori-challenge-transactions"),
	}
	_, err = svc.PutItem(input)
	return err
}

func QueryTransactionsByEmail(svc *dynamodb.DynamoDB, email string) ([]Transaction, error) {
	input := &dynamodb.QueryInput{
		TableName:              aws.String("stori-challenge-transactions"),
		KeyConditionExpression: aws.String("email = :email"),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":email": {
				S: aws.String(email),
			},
		},
	}
	result, err := svc.Query(input)
	if err != nil {
		return nil, err
	}

	var transactions []Transaction
	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &transactions)
	return transactions, err
}
