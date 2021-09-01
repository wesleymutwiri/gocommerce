package controllers

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/wesleymutwiri/gocommerce/api/responses"
)

func (server *Server) Home(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	responses.JSON(w, http.StatusOK, "Welcome")
}
