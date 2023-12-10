package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var db *sql.DB

func Connect() error {
	const (
		host     = "localhost"
		port     = 5432
		user     = "postgres"
		password = "12345"
		dbname   = "Plataform"
		sslmode  = "disable"
	)

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		host, port, user, password, dbname, sslmode)

	var err error
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		return err
	}

	// Teste a conexão
	err = db.Ping()
	if err != nil {
		return err
	}

	fmt.Println("Conectado com sucesso ao banco de dados!")
	return nil
}

// GetDB retorna a instância de conexão do banco de dados.
func GetDB() *sql.DB {
	return db
}
