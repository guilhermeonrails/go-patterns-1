package routes

import (
	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
	r := mux.NewRouter()

	ItemRoutes(r)
	CategoriaRoutes(r)

	return r
}
