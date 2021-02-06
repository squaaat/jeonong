package main

import (
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/aws/aws-lambda-go/lambda"
	adapter "github.com/awslabs/aws-lambda-go-api-proxy/fiber"
	"github.com/gofiber/fiber/v2"
)

func main() {

	f := fiber.New()

	f.Get("/swagger/*", swagger.Handler)

	lambdaApp := adapter.New(f)

	lambda.Start(lambdaApp.Proxy)
}
