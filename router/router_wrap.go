package router

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sushi-backend/types/responses"
	"sushi-backend/utils"
)

type wrappedFn func(w http.ResponseWriter, r *http.Request) *responses.Response

func (router *Router) wrapResponse(fn wrappedFn) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		resp := fn(w, r)

		if resp.IsError() {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(resp.Status)
			utils.PanicIfErrorWithResult(w.Write([]byte(fmt.Sprintf(`{"message": "%s"}`, resp.Msg))))
			return
		}

		if resp.Data != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(resp.Status)
			utils.PanicIfError(json.NewEncoder(w).Encode(&resp.Data))
			return
		}

		w.WriteHeader(resp.Status)
	}
}
