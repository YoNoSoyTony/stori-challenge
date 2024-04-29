#!/bin/bash

# Setup Node.js environment
# This step is typically handled by the GitHub Actions runner, so you might need to ensure Node.js is installed on your machine
# and use the desired version.

# Install dependencies
cd frontend
npm install

# Build frontend
npm run build

# Deploy to S3
# Make sure you have AWS CLI installed and configured
aws s3 sync ./dist/ s3://stori-challenge-frontend --delete