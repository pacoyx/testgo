package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/pacoyx/go-cron-test/controller"
)

var puertoInicio string

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	puertoInicio = os.Getenv("PUERTO_INI")
}

func main() {
	router := mux.NewRouter()

	fmt.Println("Starting the application...")

	router.HandleFunc("/authenticate", controller.CreateTokenEndpoint).Methods("POST")
	router.HandleFunc("/protected", controller.ProtectedEndpoint).Methods("GET")
	//router.HandleFunc("/test", controller.ValidateMiddleware(TestEndpoint)).Methods("GET")
	router.HandleFunc("/Pdf", controller.GenerarPDF).Methods("GET")
	router.HandleFunc("/TestMysql", controller.TestCnxMySQLController).Methods("GET")
	router.HandleFunc("/TestMysqlSelect", controller.TestCnxMySQLControllerListado).Methods("GET")
	router.HandleFunc("/TestSqlServer", controller.TestCnxSQLServerController).Methods("GET")
	router.HandleFunc("/TestSqlServerInsert", controller.TestCnxSQLServerInsertController).Methods("POST")
	router.HandleFunc("/TestSqlServerSelect", controller.TestCnxSQLServerSelectController).Methods("GET")
	router.HandleFunc("/TestMongodb", controller.TestCnxMongodbController).Methods("GET")
	router.HandleFunc("/TestMongodbInsert", controller.TestCnxMongodbInsertController).Methods("POST")
	router.HandleFunc("/TestMongodbSelect", controller.TestCnxMongodbSelectController).Methods("GET")

	log.Fatal(http.ListenAndServe(":"+puertoInicio, router))

}
