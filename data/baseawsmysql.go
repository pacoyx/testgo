package data

import (
	"fmt"
	"log"
	"os"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var clientdb *sql.DB

//GetClientDBMySQLaws xxxx
func GetClientDBMySQLaws() *sql.DB {
	initConexion()
	return clientdb
}

func initConexion() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	usu := os.Getenv("aurora_userdb")
	pwd := os.Getenv("aurora_password")
	url := os.Getenv("aurora_server")
	database := os.Getenv("aurora_database")

	db, err := sql.Open("mysql", usu+":"+pwd+"@tcp("+url+")/"+database)
	if err != nil {
		log.Fatal("Error al conectar a mysql aws")
	}

	fmt.Println("conectado a mysql aws")

	clientdb = db

}
