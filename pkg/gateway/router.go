package gateway

import (
	"github.com/gorilla/mux"
	"github.com/jcr04/Plataform-API.go/pkg/middleware"
)

// NewRouter cria e retorna um novo roteador com as rotas configuradas.
func NewRouter() *mux.Router {
	router := mux.NewRouter()

	protectedRoutes := router.PathPrefix("/protected").Subrouter()
	protectedRoutes.Use(middleware.AuthMiddleware)

	// Configuração das rotas
	protectedRoutes.HandleFunc("/example", ExampleProtectedHandler).Methods("GET")
	router.HandleFunc("/example", ExampleHandler).Methods("GET")
	// Adicione mais rotas aqui conforme necessário

	return router
}
