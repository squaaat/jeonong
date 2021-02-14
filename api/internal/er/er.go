package er

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"

	"github.com/pkg/errors"
)

type Kind int

const (
	KindUndefined Kind = iota
	KindBadRequest
	KindInternalServerError
	KindForbidden
	KindIgnorable
	KindFailJSONMarshaling
	KindFailDSLSerialize
)

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
