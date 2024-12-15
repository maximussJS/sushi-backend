package utils

import (
	"errors"
	"gorm.io/gorm"
)

func HandleRecordNotFound[T any](data T, err error) (T, error) {
	if err != nil {
		var t T
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return t, nil
		} else {
			return t, err
		}
	}

	return data, nil
}
