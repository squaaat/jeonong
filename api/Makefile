.PHONY: build clean deploy

deploy:
	# deploy server
	rm -rf ./bin
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/lambda lambda/main.go
	SLS_DEBUG=* sls deploy --verbose

logs:
	aws logs tail /aws/lambda/nearsfeed-api-alpha-api --follow