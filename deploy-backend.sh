#!/bin/bash

# Set up Go environment
export GOOS=linux
export GOARCH=amd64
export GO111MODULE=on

# Build and package handleCSV
cd backend/handleCSV
go build -o bootstrap
zip handleCSV.zip bootstrap
cd ../..

# Build and package sendEmail
cd backend/sendEmail
go build -o bootstrap
zip sendEmail.zip bootstrap
cd ../..

# Deploy to AWS Lambda
# Make sure you have AWS CLI installed and configured
aws lambda update-function-code --function-name handleCSV --zip-file fileb://backend/handleCSV/handleCSV.zip
aws lambda update-function-code --function-name sendEmail --zip-file fileb://backend/sendEmail/sendEmail.zip