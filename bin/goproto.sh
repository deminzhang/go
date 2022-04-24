#!/bin/sh

# export GOPATH=/data/app/slgdev/go
# export PATH=$PATH:/data/app/slgdev/go/bin

protoc --version
protoc --go_out=src/protos protos3/*.proto --proto_path=protos3

echo "Done!"

@pause