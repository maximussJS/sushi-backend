package utils

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"sushi-backend/types/responses"
)

func GetIdParam(r *http.Request) (string, *responses.Response) {
	id := mux.Vars(r)["id"]

	if IsValidUUID(id) == false {
		return "", responses.NewBadRequestResponse(fmt.Sprintf("Invalid id format %s. Should be UUID/V4", id))
	}

	return id, nil
}

func GetLimitQueryParam(r *http.Request, defaultLimit int) (int, *responses.Response) {
	limit := r.URL.Query().Get("limit")

	if limit == "" {
		return defaultLimit, nil
	}

	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		return 0, responses.NewBadRequestResponse(fmt.Sprintf("Invalid limit format %s. Should be integer", limit))
	}

	return limitInt, nil
}

func GetOffsetQueryParam(r *http.Request, defaultOffset int) (int, *responses.Response) {
	offset := r.URL.Query().Get("offset")

	if offset == "" {
		return defaultOffset, nil
	}

	offsetInt, err := strconv.Atoi(offset)
	if err != nil {
		return 0, responses.NewBadRequestResponse(fmt.Sprintf("Invalid offset format %s. Should be integer", offset))
	}

	return offsetInt, nil
}
