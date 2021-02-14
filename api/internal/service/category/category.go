package category

import (
	"github.com/gofiber/fiber/v2"
	jsoniter "github.com/json-iterator/go"
	"gorm.io/gorm"

	"github.com/squaaat/jeonong/api/internal/app"
	"github.com/squaaat/jeonong/api/internal/er"
	"github.com/squaaat/jeonong/api/internal/model"
	categoryStore "github.com/squaaat/jeonong/api/internal/service/category/store"
	keywordStore "github.com/squaaat/jeonong/api/internal/service/keyword/store"
)

type Service struct {
	App *app.Application
	CategoryStore *categoryStore.Service
}

func New(a *app.Application) *Service {
	return &Service{
		App: a,
		CategoryStore: categoryStore.New(a),
	}
}

type In struct {
	Categories []string `json:"categories"`
}

type Out struct {
	Categories []*model.Category `json:"categories"`
}

func (s *Service) FiberHandlerPutCategory(ctx *fiber.Ctx) error {
	op := er.CallerOp()
	ctx.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSONCharsetUTF8)

	in := new(In)
	if err := ctx.BodyParser(in); err != nil {
		err = er.WrapKindAndOp(err, er.KindBadRequest, op)
		return ctx.Status(er.ToHTTPStatus(err)).SendString(er.ToJSON(err))
	}

	out, err := s.PutCategory(in.Categories)
	if err != nil {
		err = er.WrapKindAndOp(err, er.KindInternalServerError, op)
		return ctx.Status(er.ToHTTPStatus(err)).SendString(er.ToJSON(err))
	}

	b, err := jsoniter.Marshal(out)
	if err != nil {
		err = er.WrapKindAndOp(err, er.KindInternalServerError, op)
		return ctx.Status(er.ToHTTPStatus(err)).SendString(er.ToJSON(err))
	}
	ctx.Status(fiber.StatusOK)
	return ctx.Send(b)
}

func (s *Service) PutCategory(in []string) (*Out, error) {
	op := er.CallerOp()
	var keywords = make([]*model.Keyword, len(in))
	var categories = make([]*model.Category, len(in))

	err := s.App.ServiceDB.DB.Transaction(func (tx *gorm.DB) error {
		for i, keyword := range in {
			k, err := keywordStore.MustGetKeyword(tx, keyword, "")
			if err != nil {
				return err
			}
			keywords[i] = k
		}

		for i, _ := range keywords {
			if i == (len(keywords)-1) {
				break
			}
			if i == 0 {
				cat, err := categoryStore.AddCategoryIfNotExist(tx, keywords[i], keywords[i])
				if err != nil {
					return err
				}
				categories[i] = cat
			}
			cat, err := categoryStore.AddCategoryIfNotExist(tx, keywords[i+1], keywords[i])
			if err != nil {
				return err
			}
			categories[i+1] = cat
		}
		return nil
	})
	if err != nil {
		err := er.WrapOp(err, op)
		return nil, err
	}

	return &Out{
		Categories: categories,
	}, nil
}
