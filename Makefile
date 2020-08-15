export GO111MODULE=on
export GOOS=linux

ifdef STAGE
else
STAGE:=dev
endif

.PHONY: build deploy dryrun test

build:
	go mod tidy
	go build -ldflags="-s -w" -o bin/talk talk/main.go

deploy: build
	serverless deploy --verbose --stage=${STAGE}

dryrun: build
	serverless deploy --verbose --noDeploy --stage=${STAGE}

test:
	go test ./talk
