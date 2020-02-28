package controller

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/pacoyx/go-cron-test/models"
	"github.com/pacoyx/go-cron-test/utils"
)

//TestCnxMySQLController testiando conexion a mysql aws
func TestCnxMySQLController(w http.ResponseWriter, r *http.Request) {
	mensajeMysql := models.TestCnxMysql()
	utils.RespondWithJSON(w, http.StatusCreated, mensajeMysql+" respondiendo desde controlador.... mysql ready!!")
}

//TestCnxMySQLControllerListado listado usando select
func TestCnxMySQLControllerListado(w http.ResponseWriter, r *http.Request) {
	mensajeMysql, arrResp := models.TestCnxMysqlSelect()
	log.Println(mensajeMysql)
	utils.RespondWithJSON(w, http.StatusCreated, *arrResp)
}

//TestCnxSQLServerController testiando conexion a sql server server nube
func TestCnxSQLServerController(w http.ResponseWriter, r *http.Request) {
	mensajeMSSS := models.TestCnxSQLServer()
	utils.RespondWithJSON(w, http.StatusCreated, mensajeMSSS)
}

//TestCnxSQLServerInsertController testiando conexion a sql server server nube
func TestCnxSQLServerInsertController(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "invalid request")
	}

	var eproyecto models.Proyecto
	json.Unmarshal(reqBody, &eproyecto)

	mensajeMSSS := models.TestCnxSQLServerInsert(eproyecto)
	utils.RespondWithJSON(w, http.StatusCreated, mensajeMSSS)
}

//TestCnxSQLServerSelectController listado usando select
func TestCnxSQLServerSelectController(w http.ResponseWriter, r *http.Request) {
	mensajeMSSS, arrResp := models.TestCnxSQLServerSelect()
	log.Println(mensajeMSSS)
	utils.RespondWithJSON(w, http.StatusCreated, *arrResp)
}

//TestCnxMongodbController testiando conexion a mongodb atlas
func TestCnxMongodbController(w http.ResponseWriter, r *http.Request) {
	mensajeMysql := models.TestCnxMongodb()
	utils.RespondWithJSON(w, http.StatusCreated, mensajeMysql)
}

//TestCnxMongodbInsertController test insert en mongodb
func TestCnxMongodbInsertController(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "invalid request")
	}

	var shopping models.Shopping
	json.Unmarshal(reqBody, &shopping)
	mensajeMysql := models.TestCnxMongodbInsert(shopping)
	utils.RespondWithJSON(w, http.StatusCreated, mensajeMysql)
}

//TestCnxMongodbSelectController test de select mongoDB
func TestCnxMongodbSelectController(w http.ResponseWriter, r *http.Request) {
	mensajeMysql, arrResp := models.TestCnxMongodbSelect()
	log.Println(mensajeMysql)
	utils.RespondWithJSON(w, http.StatusCreated, arrResp)
}
