package utils

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
