package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

//Message envia mensaje
func Message(status bool, message string) map[string]interface{} {
	return map[string]interface{}{"status": status, "message": message}
}

//Respond envia respuesta
func Respond(w http.ResponseWriter, data map[string]interface{}) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

//RespondWithError respuesta con formato de error
func RespondWithError(w http.ResponseWriter, code int, message string) {
	RespondWithJSON(w, code, map[string]string{"error": message})
}

//RespondWithJSON respuesta en formato JSON
func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

//GetFechaL retorna fecha en formato largo|
func GetFechaL() string {
	t := time.Now()
	fecha := fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())

	fmt.Println("La fecha larga actual es =>", fecha)
	return fecha
}

//GetFechaC retorna fecha en formato largo|
func GetFechaC() string {
	t := time.Now()
	fecha := fmt.Sprintf("%d-%02d-%02d",
		t.Year(), t.Month(), t.Day())

	fmt.Println("La fecha corta actual es =>", fecha)
	return fecha
}
