#!/bin/bash
echo "consume your new API with an HTTP/1.1 POST with a JSON payload"
curl     --header "Content-Type: application/json"     --data '{"name": "Jane"}'     http://localhost:8080/greet.v1.GreetService/Greet
echo "you should receive {\"greeting\": \"Hello, Jane!\"}"
echo "your new API supports gRPC requests, too:"
buf curl   --schema ./greet/v1/greet.proto   --protocol grpc   --http2-prior-knowledge   --data '{"name": "Jane"}'   http://localhost:8080/greet.v1.GreetService/Greet
