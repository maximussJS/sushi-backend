package utils

import (
	"errors"
	"gorm.io/gorm"
)

func HandleRecordNotFound[T any](data T, err error) T {
	if err != nil {
		var t T
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return t
		} else {
			panic(err)
		}
	}

	return data
}
