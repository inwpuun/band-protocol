package errorutil

import (
	"github.com/joomcode/errorx"
)

func WithStack(err error) error {
	if err == nil {
		return nil
	}
	return errorx.EnsureStackTrace(err)
}

func Must[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}
