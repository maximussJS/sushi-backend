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

func PanicIfErrorWithResultReturning(data interface{}, err error) (interface{}, error) {
	if err != nil {
		panic(err)
	}

	return data, nil
}
