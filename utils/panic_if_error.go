package utils

import (
	"context"
	"errors"
)

func PanicIfError(err error) {
	if err != nil {
		panic(err)
	}
}

func PanicIfErrorWithResult(_ interface{}, err error) {
	if err != nil {
		panic(err)
	}
}

func PanicIfErrorWithResultReturning[T any](data T, err error) T {
	if err != nil {
		panic(err)
	}

	return data
}

func PanicIfErrorIsNotContextError(err error) {
	if err != nil && !errors.Is(err, context.Canceled) && !errors.Is(err, context.DeadlineExceeded) {
		panic(err)
	}
}

func PanicIfErrorIsNotContextErrorWithResult(_ interface{}, err error) {
	if err != nil && !errors.Is(err, context.Canceled) && !errors.Is(err, context.DeadlineExceeded) {
		panic(err)
	}
}

func PanicIfErrorIsNotContextErrorWithResultReturning[T any](data T, err error) T {
	if err != nil && !errors.Is(err, context.Canceled) && !errors.Is(err, context.DeadlineExceeded) {
		panic(err)
	}

	return data
}
