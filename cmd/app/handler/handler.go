package handler

import (
	"cafe_role/internal/controller"
	"github.com/gorilla/mux"
	"net/http"
)

type Handler struct {
	c controller.Controller
}

func NewHandler(c controller.Controller) http.Handler {
	m := mux.NewRouter()
	//h := Handler{c: c}

	return m
}
