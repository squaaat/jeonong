package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/rs/zerolog/log"

	"github.com/squaaat/nearsfeed/api/internal/app"
	"github.com/squaaat/nearsfeed/api/internal/er"
	catSrv "github.com/squaaat/nearsfeed/api/internal/service/category"
	manSrv "github.com/squaaat/nearsfeed/api/internal/service/manufacture"
)

func New(a *app.Application) *fiber.App {
	f := fiber.New()

	f.Use(cors.New())
	f.Use(func(ctx *fiber.Ctx) error {
		if !(ctx.Get(fiber.HeaderContentType, "") == fiber.MIMEApplicationJSON || ctx.Get(fiber.HeaderContentType, "") == fiber.MIMEApplicationJSONCharsetUTF8) {
			log.Debug().
				Str("req_path_params", ctx.Params("id", "")).
				Str("req_path", ctx.Path()).
				Str("req_query", ctx.Request().URI().QueryArgs().String()).
				Str("req_header", ctx.Request().Header.String()).
				Str("req_header_bytes", string(ctx.Request().Header.Header())).
				Bytes("req_body", ctx.Body()).
				Str(fiber.HeaderContentType, ctx.Get(fiber.HeaderContentType, "")).
				Send()
			err := er.New("only accept [Content-Type: application/json]", er.KindBadRequest, "root - middleware")
			return ctx.Status(er.ToHTTPStatus(err)).SendString(er.ToJSON(err))
		}

		err := ctx.Next()
		log.Debug().
			Str("req_path_params", ctx.Params("id", "")).
			Str("req_path", ctx.Path()).
			Str("req_query", ctx.Request().URI().QueryArgs().String()).
			Str("req_header", ctx.Request().Header.String()).
			Str("req_header_bytes", string(ctx.Request().Header.Header())).
			Bytes("req_body", ctx.Body()).
			Bytes("res_body", ctx.Response().Body()).
			Str("res_header", ctx.Response().Header.String()).
			Err(err).
			Send()
		return nil
	})

	categoryService := catSrv.New(a)
	f.Get("/health", func(ctx *fiber.Ctx) error {
		return ctx.Status(fiber.StatusOK).SendString("ok")
	})
	f.Put("/api/categories", categoryService.FiberHandlerPutCategory)
	f.Get("/api/categories", categoryService.FiberHandlerGetCategories)
	f.Get("/api/categories/:id", categoryService.FiberHandlerGetCategory)
	manufactureService := manSrv.New(a)
	f.Put("/api/manufactures", manufactureService.FiberHandlerPutManufacture)
	f.Get("/api/manufactures", manufactureService.FiberHandlerGetManufactures)
	f.Get("/api/manufactures/:id", manufactureService.FiberHandlerGetManufacture)

	return f
}
