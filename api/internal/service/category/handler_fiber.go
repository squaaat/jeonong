package category

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	jsoniter "github.com/json-iterator/go"

	"github.com/squaaat/nearsfeed/api/internal/er"
)

func (s *Service) FiberHandlerPutCategory(ctx *fiber.Ctx) error {
	op := er.CallerOp()
	ctx.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSONCharsetUTF8)

	in := new(In)
	if err := ctx.BodyParser(in); err != nil {
		err = er.WrapKindAndOp(err, er.KindBadRequest, op)
		return ctx.Status(er.ToHTTPStatus(err)).SendString(er.ToJSON(err))
	}
	fmt.Println(string(ctx.Body()))

	out, err := s.PutCategory(in.Category)
	if err != nil {
		err = er.WrapOp(err, op)
		err = er.WrapKindIfNotSet(err, er.KindInternalServerError)
		return ctx.Status(er.ToHTTPStatus(err)).SendString(er.ToJSON(err))
	}

	b, err := jsoniter.Marshal(&out)
	if err != nil {
		err = er.WrapKindAndOp(err, er.KindFailJSONMarshaling, op)
		return ctx.Status(er.ToHTTPStatus(err)).SendString(er.ToJSON(err))
	}
	ctx.Status(fiber.StatusOK)
	return ctx.Send(b)
}

func (s *Service) FiberHandlerGetCategories(ctx *fiber.Ctx) error {
	op := er.CallerOp()
	ctx.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSONCharsetUTF8)

	out, err := s.GetCategories()
	if err != nil {
		err = er.WrapKindIfNotSet(er.WrapOp(err, op), er.KindInternalServerError)
		return ctx.Status(er.ToHTTPStatus(err)).SendString(er.ToJSON(err))
	}

	b, err := jsoniter.Marshal(&out)
	if err != nil {
		err = er.WrapKindAndOp(err, er.KindFailJSONMarshaling, op)
		return ctx.Status(er.ToHTTPStatus(err)).SendString(er.ToJSON(err))
	}

	ctx.Status(fiber.StatusOK)
	return ctx.Send(b)
}

func (s *Service) FiberHandlerGetCategory(ctx *fiber.Ctx) error {
	op := er.CallerOp()
	ctx.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSONCharsetUTF8)

	id := ctx.Params("id")
	if id == "" {
		err := er.New("[id] path parameter is empty", er.KindBadRequest, op)
		return ctx.Status(er.ToHTTPStatus(err)).SendString(er.ToJSON(err))
	}
	out, err := s.GetCategory(id)
	if err != nil {
		err = er.WrapKindIfNotSet(er.WrapOp(err, op), er.KindInternalServerError)
		return ctx.Status(er.ToHTTPStatus(err)).SendString(er.ToJSON(err))
	}

	b, err := jsoniter.Marshal(&out)
	if err != nil {
		err = er.WrapKindAndOp(err, er.KindFailJSONMarshaling, op)
		return ctx.Status(er.ToHTTPStatus(err)).SendString(er.ToJSON(err))
	}

	ctx.Status(fiber.StatusOK)
	return ctx.Send(b)
}
