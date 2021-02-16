package main

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	adapter "github.com/awslabs/aws-lambda-go-api-proxy/fiber"

	"github.com/squaaat/nearsfeed/api/internal/app"
	"github.com/squaaat/nearsfeed/api/internal/config"
	serverhttp "github.com/squaaat/nearsfeed/api/internal/server/http"
)


func main() {
	fmt.Println("1-----------------!!!!!!!!!!1")
	fmt.Println(os.Getenv("J_ENV"))
	cfg := config.MustInit(os.Getenv("J_ENV"), false)
	fmt.Println("2-----------------!!!!!!!!!!1")
	app := app.New(cfg)
	fmt.Println("3-----------------!!!!!!!!!!1")
	http := serverhttp.New(app)
	fmt.Println("4-----------------!!!!!!!!!!1")

	lambdaApp := adapter.New(http)
	fmt.Println("5-----------------!!!!!!!!!!1")

	lambda.Start(func (ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
		fmt.Println("6-----------------!!!!!!!!!!1")
		return lambdaApp.ProxyWithContext(ctx, req)
	})
}
