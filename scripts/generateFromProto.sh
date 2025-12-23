#!/bin/bash
echo "will generate code in gen from the proto file"
echo "📣 about to run buf dep update"
buf dep update
echo "📣 about to run buf lint"
buf lint
echo "📣 about to run buf generate"
buf generate
