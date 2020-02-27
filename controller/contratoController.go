package controller

import (
	"net/http"

	"github.com/pacoyx/go-cron-test/utils"
)

//TestController validamos la conexion
func TestController(w http.ResponseWriter, r *http.Request) {

	utils.RespondWithJSON(w, http.StatusCreated, "conectado")
}
