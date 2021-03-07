package http

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/rs/zerolog/log"

	_const "github.com/squaaat/nearsfeed/api/internal/const"
	"github.com/squaaat/nearsfeed/api/internal/container"
	catSrv "github.com/squaaat/nearsfeed/api/internal/service/category"
	manSrv "github.com/squaaat/nearsfeed/api/internal/service/manufacture"
)

func ConfigCORS(origin, env string) *cors.Config {
	cfg := &cors.ConfigDefault

	if env == _const.EnvAlpha {
		cfg.AllowCredentials = false
		cfg.AllowOrigins = origin
	}
	return cfg
}

func New(cont *container.Container) *fiber.App {
	f := fiber.New()

	f.Use(func(ctx *fiber.Ctx) error {
		reqOrigin := ctx.Get(fiber.HeaderOrigin, "*")
		allowMethods := "GET,POST,HEAD,PUT,DELETE,PATCH"
		ctx.Set(fiber.HeaderAccessControlAllowMethods, allowMethods)


		fmt.Println(reqOrigin)
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
	f.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))
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
