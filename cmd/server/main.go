package main

import (
	"log"
	"net/http"

	"github.com/jcr04/Plataform-API.go/pkg/gateway"
	"github.com/jcr04/Plataform-API.go/pkg/utils"
)

func main() {
	utils.InfoLogger.Println("Iniciando a aplicação...")
	router := gateway.NewRouter()

	// Configuração do servidor HTTP
	http.Handle("/", router)
	log.Println("Starting the server on :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
	utils.ErrorLogger.Println("Erro ao processar a solicitação...")
}
