# stori-challenge

This project is a comprehensive solution for managing and processing data, including CSV files and sending emails. It consists of a backend and a frontend, each with its own CI/CD pipeline.

Test it out [here](https://stori-challenge-frontend.s3.us-east-2.amazonaws.com/index.html)! 
There is a testfile called file.csv at the root of this project


## Prerequisites

Before deploying this project, ensure the following AWS resources are already set up:

1. **AWS Lambda Functions**: Two Lambda functions named `handleCSV` and `sendEmail` must be created. these functions need to be configured to be triggered by an Amazon API Gateway.

2. **Amazon S3 Bucket**: An S3 bucket named `stori-challenge-frontend` must be created. This bucket will host the frontend application's static files.

## Deployment

The project uses GitHub Actions for CI/CD, with separate workflows for the backend and frontend. The backend workflow builds Go applications and deploys them to AWS Lambda, while the frontend workflow builds the frontend application and uploads the build artifacts to the S3 bucket.

## Future Improvements

While the current setup works well, it's recommended to consider using AWS Serverless Application Model (SAM) for deploying the backend and frontend in the future. AWS SAM simplifies the process of building, deploying, and managing serverless applications on AWS. It provides a framework for defining serverless resources and their properties in a simple and concise manner.

## Locally
To deploy this project locally, you will need the same prerequisites as mentioned above, including the AWS Lambda functions and the S3 bucket. Additionally, ensure you have the AWS CLI set up on your machine.

To deploy the frontend and backend, navigate to the root of the repository and run the following scripts:

- For the frontend: `./deployFrontend.sh`
- For the backend: `./deployBackend.sh`
