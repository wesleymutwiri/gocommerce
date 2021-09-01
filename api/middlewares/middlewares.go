package middlewares

import (
	"errors"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/wesleymutwiri/gocommerce/api/auth"
	"github.com/wesleymutwiri/gocommerce/api/responses"
)

func SetMiddlewareJSON(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")
		next(w, r, ps)
	}
}

func SetMiddlewareAuthentication(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		err := auth.TokenValid(r)
		if err != nil {
			responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
			return
		}
		next(w, r, ps)
	}
}
