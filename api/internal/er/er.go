package er

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"

	"github.com/gofiber/fiber/v2"
	jsoniter "github.com/json-iterator/go"
	"github.com/pkg/errors"
)

type Kind int

const (
	KindUndefined Kind = iota
	KindBadRequest
	KindNotFound
	KindInternalServerError
	KindForbidden
	KindIgnorable
	KindFailJSONMarshaling
	KindDubplicated
)

var (
	KindAndHTTPStatusMap = map[Kind]int{
		KindUndefined:           fiber.StatusInternalServerError,
		KindBadRequest:          fiber.StatusBadRequest,
		KindNotFound:            fiber.StatusNotFound,
		KindInternalServerError: fiber.StatusInternalServerError,
		KindFailJSONMarshaling:  fiber.StatusInternalServerError,
		KindDubplicated:         fiber.StatusBadRequest,
		KindForbidden:           fiber.StatusForbidden,
	}
)

func (k Kind) String() string {
	switch k {
	case KindUndefined:
		return "KindUndefined"
	case KindBadRequest:
		return "KindBadRequest"
	case KindInternalServerError:
		return "KindInternalServerError"
	case KindForbidden:
		return "KindForbidden"
	case KindIgnorable:
		return "KindIgnorable"
	case KindFailJSONMarshaling:
		return "KindFailJSONMarshaling"
	default:
		return "KindUndefined"
	}
}

type Error struct {
	Ops  []string
	Kind Kind
	Err  error
}

func New(errMsg string, kind Kind, op string) error {
	return WrapKindAndOp(errors.New(errMsg), kind, op)
}

func new(err error) *Error {
	if e, ok := err.(*Error); ok {
		return e
	}
	return &Error{Err: err}
}

func Is(sourceErr, targetErr error) bool {
	se := new(sourceErr)
	te := new(targetErr)

	return errors.Is(se.Err, te.Err)
}

func WrapOp(err error, op string) error {
	e := new(err)
	e.Ops = append(e.Ops, op)
	return e
}

func WrapKindIfNotSet(err error, kind Kind) error {
	e := new(err)
	if e.Kind != KindUndefined {
		return e
	}
	e.Kind = kind
	return e
}

func WrapKind(err error, kind Kind) error {
	e := new(err)
	e.Kind = kind
	return e
}

func IsKind(err error, kind Kind) bool {
	e := new(err)
	return e.Kind == kind
}

func WrapKindAndOp(err error, kind Kind, op string) error {
	e := new(err)
	e.Ops = append(e.Ops, op)
	e.Kind = kind
	return e
}

func (e *Error) Error() string {
	if e.Err == nil {
		e.Err = errors.New("")
	}
	et := reflect.TypeOf(e.Err)
	originMsg := fmt.Sprintf("%s[%s]", e.Err.Error(), et.String())
	ops := []string{originMsg}
	ops = append(ops, e.Ops...)
	return strings.Join(ops, "/")
}

func CallerOp() string {
	pc, _, _, _ := runtime.Caller(1)
	caller := runtime.FuncForPC(pc).Name()
	splits := strings.Split(caller, "/")
	return strings.Join(splits[3:], ".")
}

type ErrorJSON struct {
	Error string `json:"error"`
	Kind  string `json:"kind"`
}

func ToJSON(err error) string {
	e := new(err)
	jsonObj := ErrorJSON{
		Error: e.Error(),
		Kind:  e.Kind.String(),
	}
	s, err := jsoniter.MarshalToString(&jsonObj)
	if err != nil {
		return `{"error": "er.(e *Error).ToJSON marshalling failed", "kind": "KindFailJSONMarshaling"}`
	}
	return s
}

func ToHTTPStatus(err error) int {
	e := new(err)
	if i, ok := KindAndHTTPStatusMap[e.Kind]; ok {
		return i
	}
	return fiber.StatusInternalServerError
}
