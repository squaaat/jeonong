package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/rs/zerolog/log"

	"github.com/squaaat/nearsfeed/api/internal/container"
	catSrv "github.com/squaaat/nearsfeed/api/internal/service/category"
	manSrv "github.com/squaaat/nearsfeed/api/internal/service/manufacture"
)

func New(cont *container.Container) *fiber.App {
	f := fiber.New()

	f.Use(cors.New())
	f.Use(func(ctx *fiber.Ctx) error {
		err := ctx.Next()
		log.Debug().
			Str("req_path_params", ctx.Params("id", "")).
			Str("req_uri", ctx.Path()+"?"+ctx.Request().URI().QueryArgs().String()).
			Str("req_header", ctx.Request().Header.String()).
			Bytes("req_body", ctx.Body()).
			Str("res_header", ctx.Response().Header.String()).
			Bytes("res_body", ctx.Response().Body()).
			Err(err).
			Send()
		return nil
	})

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
