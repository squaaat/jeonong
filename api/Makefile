.PHONY: build clean deploy

run:
	# locally run
	go mod download
	go run main.go http start

swag:
	# create swagger.yml
	swagger generate spec -o ./swagger.yml --scan-models
    aws s3 cp ./swagger.yml s3://squaaat-lambda/serverless/jeonong-api/alpha/swagger.yml --acl public-read

swag-val:
	# validate swagger.yml
	swagger validate ./swagger.yml

deploy: swag swag-val
	# deploy server
	rm -rf ./bin
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/lambda lambda/main.go
	SLS_DEBUG=* sls deploy --verbose

gorm-create:
	go run main.go gorm create

gorm-clean:
	go run main.go gorm clean

gorm-recreate:
	go run main.go gorm re-create

gorm-migrate-create:
	go run main.go gorm migrate create

gorm-migrate-sync:
	go run main.go gorm migrate sync