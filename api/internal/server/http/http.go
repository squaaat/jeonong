package http

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/squaaat/jeonong/api/internal/app"
	catSrv "github.com/squaaat/jeonong/api/internal/service/category"
)

func New(a *app.Application) *fiber.App {
	f := fiber.New()

	f.Use(cors.New())
	f.Use(func(ctx *fiber.Ctx) error {
		fmt.Println(ctx.Path())
		fmt.Println(string(ctx.Body()))
		return ctx.Next()
	})

	categoryService := catSrv.New(a)
	f.Put("/api/categories", categoryService.FiberHandlerPutCategory)
	f.Get("/api/categories", categoryService.FiberHandlerGetCategories)
	return f
}
