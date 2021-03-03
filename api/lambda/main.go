package main

import (
	"context"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/rs/zerolog/log"

	adapter "github.com/squaaat/nearsfeed/api/pkg/lambdaadapter"

	"github.com/squaaat/nearsfeed/api/internal/app"
	"github.com/squaaat/nearsfeed/api/internal/config"
	serverhttp "github.com/squaaat/nearsfeed/api/internal/server/http"
)

func main() {
	cfg := config.MustInit(os.Getenv("J_ENV"), false)
	app := app.New(cfg)
	http := serverhttp.New(app)

	lambdaApp := adapter.New(http)

	lambda.Start(func(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
		log.Debug().Interface("API_Gateway_Proxy_header", req.Headers).Send()
		return lambdaApp.ProxyWithContext(ctx, req)
	})
}
