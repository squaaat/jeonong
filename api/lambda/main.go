package main

import (
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	adapter "github.com/awslabs/aws-lambda-go-api-proxy/fiber"

	"github.com/squaaat/nearsfeed/api/internal/app"
	"github.com/squaaat/nearsfeed/api/internal/config"
	serverhttp "github.com/squaaat/nearsfeed/api/internal/server/http"
)

func main() {
	cfg := config.MustInit(os.Getenv("J_ENV"), false)
	app := app.New(cfg)
	http := serverhttp.New(app)

	lambdaApp := adapter.New(http)

	lambda.Start(lambdaApp.Proxy)
}
