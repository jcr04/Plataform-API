package utils

import (
	"log"
	"os"
)

var (
	// InfoLogger é usado para logs informativos
	InfoLogger *log.Logger

	// ErrorLogger é usado para logs de erros
	ErrorLogger *log.Logger
)

func init() {
	// Configurando InfoLogger
	file, err := os.OpenFile("info.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	InfoLogger = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)

	// Configurando ErrorLogger
	file, err = os.OpenFile("error.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	ErrorLogger = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}
