export GO111MODULE=on
export GOOS=linux

# .PHONY: clean
# 	rm -r ./bin

build:
	go mod tidy
	go build -ldflags="-s -w" -o bin/hello hello/main.go
	go build -ldflags="-s -w" -o bin/world world/main.go

.PHONY: deploy
deploy: build
	serverless deploy --verbose

.PHONY: dryrun
deploy: build
	serverless deploy --verbose --noDeploy
