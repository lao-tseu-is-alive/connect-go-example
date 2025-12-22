#!/bin/bash
echo "will generate code in gen from greet/v1/greet.proto"
buf dep update
buf lint
buf generate
