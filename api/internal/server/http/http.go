package http

import (
	"fmt"

	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/squaaat/jeonong/api/internal/app"
	catSrv "github.com/squaaat/jeonong/api/internal/service/category"
)

// JEONONG-api application http api server that jeonong-api
//
// the purpose of this application is to provide an application
// that is using plain go code to define an API
//
//     Schemes: http, https
//     Host: localhost
//     BasePath: /api/v2
//     Version: 0.0.1
//     License: MIT http://opensource.org/licenses/MIT
//     Contact: drakejin<dydwls121200@gmail.com> https://github.com/drakejin
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
// swagger:meta

func New(a *app.Application) *fiber.App {
	f := fiber.New()

	f.Use(cors.New())
	f.Use(func(ctx *fiber.Ctx) error {
		fmt.Println(ctx.Path())
		fmt.Println(string(ctx.Body()))
		return ctx.Next()
	})

	if a.Config.App.Debug {
		//f.Get("/swagger/*", swagger.New(swagger.Config{URL: fmt.Sprintf("https://squaaat-lambda.s3.ap-northeast-2.amazonaws.com/serverless/%s/%s/swagger.yml", a.Config.App.AppName, a.Config.App.Env)}))
		f.Get("/swagger/*", swagger.Handler)
	}

	categoryService := catSrv.New(a)
	f.Put("/v1/categories", categoryService.FiberHandlerPutCategory)

	return f
}
