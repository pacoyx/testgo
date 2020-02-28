package data

//https://mathaywardhill.com/2017/04/27/get-started-with-golang-and-sql-server-in-visual-studio-code/
//https://github.com/microsoft/sql-server-samples/blob/master/samples/tutorials/go/crud.go

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/joho/godotenv"
)

var clientemmss *sql.DB

//GetClienteMMSS devuelve el cliente cnx para sql server
func GetClienteMMSS() *sql.DB {
	return clientemmss
}

//GetPruebaCnx test de coenxion
func GetPruebaCnx() string {
	return "conexion MMSS OK"
}

func init() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	usu := os.Getenv("MSSS_USER")
	pwd := os.Getenv("MSSS_PWD")
	urlServer := os.Getenv("MSSS_SERVER")
	database := os.Getenv("MSSS_DATABASE")

	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;database=%s;",
		urlServer, usu, pwd, database)

	condb, errdb := sql.Open("mssql", connString)
	if errdb != nil {
		fmt.Println(" Error open db:", errdb.Error())
	}
	var (
		sqlversion string
	)
	rows, err := condb.Query("select @@version")
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		err := rows.Scan(&sqlversion)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(sqlversion)
	}
	clientemmss = condb
	//defer condb.Close()
}
