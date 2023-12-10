package gateway

import (
	"log"
	"net/http/httputil"
	"net/url"

	"github.com/gorilla/mux"
	"github.com/jcr04/Plataform-API.go/pkg/middleware"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter()

	auauPetsURL, err := url.Parse("http://localhost:8081")
	if err != nil {
		log.Fatal(err)
	}

	auauPetsProxy := httputil.NewSingleHostReverseProxy(auauPetsURL)

	router.Handle("/animals", auauPetsProxy)
	router.Handle("/animals/{id}", auauPetsProxy)
	router.Handle("/reservations", auauPetsProxy)
	router.Handle("/reservations/{id}", auauPetsProxy)
	router.Handle("/hostings", auauPetsProxy)
	router.Handle("/hostings/{id}", auauPetsProxy)

	protectedRoutes := router.PathPrefix("/protected").Subrouter()
	protectedRoutes.Use(middleware.AuthMiddleware)

	// Configuração das rotas
	protectedRoutes.HandleFunc("/example", ExampleProtectedHandler).Methods("GET")
	router.HandleFunc("/example", ExampleHandler).Methods("GET")
	// Adicione mais rotas aqui conforme necessário

	return router
}
