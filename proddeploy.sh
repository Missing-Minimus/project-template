#!/bin/bash
echo '---'
echo timestamp=$(date +%d-%m-%Y_%H-%M-%S)
mv .env .env.tmp
cp .env.prod .env
rm -rf serverless.yaml && cp serverless-prod.yaml serverless.yaml && serverless deploy --stage production && rm -rf serverless.yaml
rm -rf .env && cp .env.tmp .env
rm -rf .env.tmp
echo '---'
