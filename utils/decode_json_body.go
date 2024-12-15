package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"io"
	"net/http"
	"strings"
	"sushi-backend/types/responses"
)

var validate = validator.New()

func DecodeJSONBody(w http.ResponseWriter, r *http.Request, dst interface{}) *responses.Response {
	ct := r.Header.Get("Content-Type")
	if ct != "" {
		mediaType := strings.ToLower(strings.TrimSpace(strings.Split(ct, ";")[0]))
		if mediaType != "application/json" {
			return responses.NewUnsupportedMediaTypeResponse("Content-Type header is not application/json")
		}
	}

	r.Body = http.MaxBytesReader(w, r.Body, 1048576)

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	err := dec.Decode(&dst)
	if err != nil {
		var syntaxError *json.SyntaxError
		var unmarshalTypeError *json.UnmarshalTypeError

		switch {
		case errors.As(err, &syntaxError):
			msg := fmt.Sprintf("Request body contains badly-formed JSON (at position %d)", syntaxError.Offset)
			return responses.NewBadRequestResponse(msg)

		case errors.Is(err, io.ErrUnexpectedEOF):
			msg := fmt.Sprintf("Request body contains badly-formed JSON")
			return responses.NewBadRequestResponse(msg)

		case errors.As(err, &unmarshalTypeError):
			msg := fmt.Sprintf("Request body contains an invalid value for the %q field (at position %d)", unmarshalTypeError.Field, unmarshalTypeError.Offset)
			return responses.NewBadRequestResponse(msg)

		case strings.HasPrefix(err.Error(), "json: unknown field "):
			fieldName := strings.TrimPrefix(err.Error(), "json: unknown field ")
			msg := fmt.Sprintf("Request body contains unknown field %s", fieldName)
			return responses.NewBadRequestResponse(msg)

		case errors.Is(err, io.EOF):
			msg := "Request body must not be empty"
			return responses.NewBadRequestResponse(msg)

		case err.Error() == "http: request body too large":
			msg := "Request body must not be larger than 1MB"
			return responses.NewBadRequestResponse(msg)

		default:
			return responses.NewInternalServerErrorResponse(err.Error())
		}
	}

	err = dec.Decode(&struct{}{})
	if !errors.Is(err, io.EOF) {
		msg := "Request body must only contain a single JSON object"
		return responses.NewBadRequestResponse(msg)
	}

	if err := validate.Struct(dst); err != nil {
		return responses.NewBadRequestResponse(err.Error())
	}

	return nil
}
