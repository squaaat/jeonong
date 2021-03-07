package http

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"github.com/squaaat/nearsfeed/api/internal/container"
	catSrv "github.com/squaaat/nearsfeed/api/internal/service/category"
	manSrv "github.com/squaaat/nearsfeed/api/internal/service/manufacture"
)

func New(cont *container.Container) *fiber.App {
	f := fiber.New()

	// CORS middleware
	f.Use(func(ctx *fiber.Ctx) error {
		reqOrigin := ctx.Get(fiber.HeaderOrigin, "*")
		allowMethods := "GET,POST,HEAD,PUT,DELETE,PATCH"
		ctx.Set(fiber.HeaderAccessControlAllowMethods, allowMethods)


		if strings.Contains(reqOrigin, "nearsfeed") {
			ctx.Set(fiber.HeaderAccessControlAllowCredentials, "true")
			ctx.Set(fiber.HeaderAccessControlAllowOrigin, reqOrigin)
		} else if strings.Contains(reqOrigin, "localhost") {
			ctx.Set(fiber.HeaderAccessControlAllowCredentials, "false")
			ctx.Set(fiber.HeaderAccessControlAllowOrigin, reqOrigin)
		} else {
			ctx.Status(fiber.StatusBadRequest)
		}
		return ctx.Next()
	})
	// req, res logger middleware
	f.Use(logger.New())

	f.Get("/health", func(ctx *fiber.Ctx) error {
		return ctx.Status(fiber.StatusOK).SendString("ok")
	})

	categoryService := catSrv.New(cont)
	f.Put("/api/categories", categoryService.FiberHandlerPutCategory)
	f.Get("/api/categories", categoryService.FiberHandlerGetCategories)
	f.Get("/api/categories/:id", categoryService.FiberHandlerGetCategory)

	manufactureService := manSrv.New(cont)
	f.Put("/api/manufactures", manufactureService.FiberHandlerPutManufacture)
	f.Get("/api/manufactures", manufactureService.FiberHandlerGetManufactures)
	f.Get("/api/manufactures/:id", manufactureService.FiberHandlerGetManufacture)

	return f
}
