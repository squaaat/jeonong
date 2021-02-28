package manufacture

import (
	"github.com/gofiber/fiber/v2"
	jsoniter "github.com/json-iterator/go"

	"github.com/squaaat/nearsfeed/api/internal/er"
)

func (s *Service) FiberHandlerPutManufacture(ctx *fiber.Ctx) error {
	op := er.CallerOp()
	ctx.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSONCharsetUTF8)

	in := new(In)
	if err := ctx.BodyParser(in); err != nil {
		err = er.WrapKindAndOp(err, er.KindBadRequest, op)
		return ctx.Status(er.ToHTTPStatus(err)).SendString(er.ToJSON(err))
	}

	out, err := s.PutManufacture(in.Manufacture)
	if err != nil {
		err = er.WrapKindAndOp(err, er.KindInternalServerError, op)
		return ctx.Status(er.ToHTTPStatus(err)).SendString(er.ToJSON(err))
	}

	b, err := jsoniter.Marshal(&out)
	if err != nil {
		err = er.WrapKindAndOp(err, er.KindInternalServerError, op)
		return ctx.Status(er.ToHTTPStatus(err)).SendString(er.ToJSON(err))
	}
	ctx.Status(fiber.StatusOK)
	return ctx.Send(b)
}

func (s *Service) FiberHandlerGetManufactures(ctx *fiber.Ctx) error {
	op := er.CallerOp()
	ctx.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSONCharsetUTF8)

	out, err := s.GetManufactures()
	if err != nil {
		err = er.WrapKindAndOp(err, er.KindInternalServerError, op)
		return ctx.Status(er.ToHTTPStatus(err)).SendString(er.ToJSON(err))
	}

	b, err := jsoniter.Marshal(&out)
	if err != nil {
		err = er.WrapKindAndOp(err, er.KindInternalServerError, op)
		return ctx.Status(er.ToHTTPStatus(err)).SendString(er.ToJSON(err))
	}

	ctx.Status(fiber.StatusOK)
	return ctx.Send(b)
}


func (s *Service) FiberHandlerGetManufacture(ctx *fiber.Ctx) error {
	op := er.CallerOp()
	ctx.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSONCharsetUTF8)

	id := ctx.Params("id")
	out, err := s.GetManufacture(id)
	if err != nil {
		err = er.WrapKindAndOp(err, er.KindInternalServerError, op)
		return ctx.Status(er.ToHTTPStatus(err)).SendString(er.ToJSON(err))
	}

	b, err := jsoniter.Marshal(&out)
	if err != nil {
		err = er.WrapKindAndOp(err, er.KindInternalServerError, op)
		return ctx.Status(er.ToHTTPStatus(err)).SendString(er.ToJSON(err))
	}

	ctx.Status(fiber.StatusOK)
	return ctx.Send(b)
}
