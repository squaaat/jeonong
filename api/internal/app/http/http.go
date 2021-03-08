package http

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"github.com/squaaat/nearsfeed/api/internal/container"
	catSrv "github.com/squaaat/nearsfeed/api/internal/service/category"
	manSrv "github.com/squaaat/nearsfeed/api/internal/service/manufacture"
)

func New(cont *container.Container) *fiber.App {
	f := fiber.New()

	// CORS middleware
	f.Use(cors.New())
	f.Use(func(ctx *fiber.Ctx) error {
		// methods
		ctx.Set(fiber.HeaderAccessControlAllowMethods, strings.Join([]string{
			fiber.MethodGet,
			fiber.MethodPost,
			fiber.MethodHead,
			fiber.MethodPut,
			fiber.MethodDelete,
			fiber.MethodPatch,
			fiber.MethodOptions,
		}, ","))
		ctx.Vary(fiber.HeaderAccessControlRequestMethod)

		ctx.Set(fiber.HeaderAccessControlAllowHeaders, strings.Join([]string{
			strings.ToLower(fiber.HeaderContentType),
		}, ", "))
		ctx.Vary(fiber.HeaderAccessControlRequestHeaders)

		reqOrigin := ctx.Get(fiber.HeaderOrigin, "*")
		if strings.Contains(reqOrigin, "nearsfeed") {
			ctx.Set(fiber.HeaderAccessControlAllowCredentials, "true")
			ctx.Set(fiber.HeaderAccessControlAllowOrigin, reqOrigin)
		} else if strings.Contains(reqOrigin, "localhost") {
			ctx.Set(fiber.HeaderAccessControlAllowCredentials, "false")
			ctx.Set(fiber.HeaderAccessControlAllowOrigin, reqOrigin)
		} else {
			return ctx.Next()
		}
		ctx.Vary(fiber.HeaderOrigin)

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
