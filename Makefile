export GO111MODULE=on
export GOOS=linux

.PHONY: build deploy dryrun test

build:
	go mod tidy
	go build -ldflags="-s -w" -o bin/talk talk/main.go

deploy: build
	serverless deploy --verbose

dryrun: build
	serverless deploy --verbose --noDeploy

test:
	go test ./talk
