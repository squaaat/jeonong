package http

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/squaaat/nearsfeed/api/internal/app"
	catSrv "github.com/squaaat/nearsfeed/api/internal/service/category"
	manSrv "github.com/squaaat/nearsfeed/api/internal/service/manufacture"
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

	manufactureService := manSrv.New(a)
	f.Put("/api/manufactures", manufactureService.FiberHandlerPutManufacture)
	f.Get("/api/manufactures", manufactureService.FiberHandlerGetManufactures)

	return f
}
