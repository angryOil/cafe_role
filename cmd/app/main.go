package main

import (
	"cafe_role/cmd/app/handler"
	"cafe_role/internal/controller"
	"cafe_role/internal/repository"
	"cafe_role/internal/repository/infla"
	"cafe_role/internal/service"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	h := getHandler()

	r.PathPrefix("/roles").Handler(h)
	http.ListenAndServe(":8086", r)
}

func getHandler() http.Handler {
	return handler.NewHandler(controller.NewController(service.NewService(repository.NewRepository(infla.NewDB()))))
}
