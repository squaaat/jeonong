package main

import (
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/rs/zerolog/log"

	"github.com/squaaat/nearsfeed/api/internal/container"
	adapter "github.com/squaaat/nearsfeed/api/pkg/lambdaadapter"

	apphttp "github.com/squaaat/nearsfeed/api/internal/app/http"
	"github.com/squaaat/nearsfeed/api/internal/config"
)

func main() {
	cfg, err := config.New(os.Getenv("J_ENV"))
	if err != nil {
		log.Fatal().Err(err).Send()
	}

	cont, err := container.New(cfg)
	if err != nil {
		log.Fatal().Err(err).Send()
	}
	http := apphttp.New(cont)

	lambdaApp := adapter.New(http)
	lambda.Start(lambdaApp.ProxyWithContext)
}
