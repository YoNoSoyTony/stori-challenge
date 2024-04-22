package shared

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

// NewDynamoDBClient creates a new DynamoDB client.
func NewDynamoDBClient() (*dynamodb.DynamoDB, error) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"),
	})
	if err != nil {
		return nil, err
	}
	return dynamodb.New(sess), nil
}

// PutTransactions uses BatchWriteItem to put multiple transactions into DynamoDB.
func PutTransactions(svc *dynamodb.DynamoDB, transactions []Transaction) error {
	requests := make([]*dynamodb.WriteRequest, len(transactions))
	for i, transaction := range transactions {
		transaction.GenerateTransactionID() // Generate a unique transaction ID
		item, err := dynamodbattribute.MarshalMap(transaction)
		if err != nil {
			return err
		}
		requests[i] = &dynamodb.WriteRequest{
			PutRequest: &dynamodb.PutRequest{
				Item: item,
			},
		}
	}

	input := &dynamodb.BatchWriteItemInput{
		RequestItems: map[string][]*dynamodb.WriteRequest{
			"stori-challenge-transactions": requests,
		},
	}

	_, err := svc.BatchWriteItem(input)
	return err
}

// QueryTransactionsByEmail queries DynamoDB for transactions by email.
func QueryTransactionsByEmail(svc dynamodbiface.DynamoDBAPI, email string) ([]Transaction, error) {
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
	if err != nil {
		return nil, err
	}

	return transactions, nil
}
