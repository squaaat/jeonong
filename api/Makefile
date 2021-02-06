.PHONY: build clean deploy

build:
	go mod download
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/lambda lambda/main.go
	cp ./swagger.yml ./bin
clean:
	rm -rf ./bin

deploy: clean build
	SLS_DEBUG=* sls deploy --verbose

