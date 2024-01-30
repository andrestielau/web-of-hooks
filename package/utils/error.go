package utils

import (
	"errors"
	"fmt"

	"github.com/samber/lo"
)

type Error struct {
	Code   int
	Index  string
	Reason string
}

func (e Error) Error() string {
	return fmt.Sprintf("error %d at %s: %s", e.Code, e.Index, e.Reason)
}

type Errors []Error

func (e Errors) Error() string   { return errors.Join(e.Unwrap()...).Error() }
func (e Errors) Unwrap() []error { return lo.Map(e, func(i Error, _ int) error { return i }) }
