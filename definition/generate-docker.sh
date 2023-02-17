#!/bin/bash
ACCOUNT_ID=$(aws sts get-caller-identity --output json | jq -r ".Account")
REGION="eu-west-1"
ECR_URL="$ACCOUNT_ID".dkr.ecr."$REGION".amazonaws.com
REPO_URL="$ECR_URL"/spotify-mock

mockoon-cli dockerize --data ./Spotify.json --port 3000 --output Dockerfile
aws ecr get-login-password --region "$REGION" | docker login --username AWS --password-stdin "$ECR_URL"
docker build -t spotify-mock  .
docker tag spotify-mock "$REPO_URL"
docker push "$REPO_URL"
