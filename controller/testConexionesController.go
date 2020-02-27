package controller

import (
	"net/http"

	"github.com/pacoyx/go-cron-test/models"
	"github.com/pacoyx/go-cron-test/utils"
)

//TestCnxMySQLController testiando conexion a mysql aws
func TestCnxMySQLController(w http.ResponseWriter, r *http.Request) {
	mensajeMysql := models.TestCnxMysql()
	utils.RespondWithJSON(w, http.StatusCreated, mensajeMysql+" responpiendo desde controlador.... mysql ready!!")
}

//TestCnxSQLServerController testiando conexion a sql server server nube
func TestCnxSQLServerController(w http.ResponseWriter, r *http.Request) {
	utils.RespondWithJSON(w, http.StatusCreated, "respondiendo desde controlador.... sql server ready!!")
}

//TestCnxMongodbController testiando conexion a mongodb atlas
func TestCnxMongodbController(w http.ResponseWriter, r *http.Request) {
	utils.RespondWithJSON(w, http.StatusCreated, "respondiendo desde controlador.... mongodb ready!!")
}
