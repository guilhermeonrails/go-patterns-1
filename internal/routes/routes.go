package routes

import (
	"myapi/internal/middleware"

	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
	r := mux.NewRouter()

	r.Use(middleware.JsonContentType)
	ItemRoutes(r)
	CategoriaRoutes(r)

	return r
}
