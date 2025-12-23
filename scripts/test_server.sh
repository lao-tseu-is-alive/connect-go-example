#!/bin/bash
# Test script for the greet server
# Demonstrates REST and RPC endpoints

set -e

echo "========================================"
echo "Testing REST Endpoints (for OpenAPI clients)"
echo "========================================"

echo ""
echo "1. GET /v1/greet/{name} - REST with path parameter"
echo "   curl http://localhost:8080/v1/greet/Alice"
curl -s http://localhost:8080/v1/greet/Alice
echo ""

echo ""
echo "2. POST /v1/greet - REST with JSON body"
echo "   curl -X POST -H 'Content-Type: application/json' -d '{\"name\":\"Bob\"}' http://localhost:8080/v1/greet"
curl -s -X POST --header "Content-Type: application/json" --data '{"name": "Bob"}' http://localhost:8080/v1/greet
echo ""

echo ""
echo "========================================"
echo "Testing RPC Endpoints (Connect/gRPC)"
echo "========================================"

echo ""
echo "3. Connect protocol via Go client"
echo "   go run ./cmd/client/client.go -mode=connect -name=Charlie"
go run ./cmd/client/client.go -mode=connect -name=Charlie 2>&1 | tail -2

echo ""
echo "4. gRPC protocol via Go client"
echo "   go run ./cmd/client/client.go -mode=grpc -name=Diana"
go run ./cmd/client/client.go -mode=grpc -name=Diana 2>&1 | tail -1

echo ""
echo "5. gRPC via buf curl (requires HTTP/2)"
echo "   buf curl --schema ./greet/v1/greet.proto --protocol grpc --http2-prior-knowledge --data '{\"name\": \"Eve\"}' http://localhost:8080/greet.v1.GreetService/Greet"
buf curl --schema ./greet/v1/greet.proto --protocol grpc --http2-prior-knowledge --data '{"name": "Eve"}' http://localhost:8080/greet.v1.GreetService/Greet

echo ""
echo "========================================"
echo "All tests completed!"
echo "========================================"
