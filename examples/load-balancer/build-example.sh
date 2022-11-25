#!/bin/bash

cd ../../

echo "Building misty broker..."

docker build --no-cache -t broker -f broker/Dockerfile .

echo "Building misty entrypoint..."

docker build --no-cache -t misty-entrypoint -f examples/load-balancer/entrypoint/Dockerfile .

echo "Building misty logger..."

docker build --no-cache -t misty-logger -f examples/load-balancer/logger/Dockerfile .

echo "Building misty sample services..."

docker build --no-cache -t misty-sample-service -f examples/load-balancer/sample-services/Dockerfile .
