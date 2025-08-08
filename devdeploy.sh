#!/bin/bash
echo '---'
echo timestamp=$(date +%d-%m-%Y_%H-%M-%S)
mv .env .env.tmp
cp .env.dev .env
rm -rf serverless.yaml && cp serverless-dev.yaml serverless.yaml && serverless deploy && rm -rf serverless.yaml
rm -rf .env && cp .env.tmp .env
rm -rf .env.tmp
echo '---'
