package utils

import (
	"encoding/json"
)

func MustJson(data interface{}) []byte {
	result, err := json.Marshal(data)
	PanicIfError(err)
	return result
}
