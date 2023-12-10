package gateway

import (
	"net/http"
)

// ExampleHandler é um exemplo de handler para rotas não protegidas
func ExampleHandler(w http.ResponseWriter, r *http.Request) {
	// Sua lógica aqui
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Este é um endpoint de exemplo!"))
}

// ExampleProtectedHandler é um exemplo de handler para rotas protegidas
func ExampleProtectedHandler(w http.ResponseWriter, r *http.Request) {
	// Sua lógica aqui
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Este é um endpoint protegido de exemplo!"))
}
