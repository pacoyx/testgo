package controller

import (
	"log"
	"net/http"

	"github.com/pacoyx/go-cron-test/utils"
)

//GenerarPDF genera pdf en disco
func GenerarPDF(w http.ResponseWriter, r *http.Request) {

	utils.CrearDirectorioSiNoExiste("./pdfs")
	data := utils.LoadCSV(utils.Path())
	pdf := utils.NewReport()

	pdf = utils.Header(pdf, data[0])
	pdf = utils.Table(pdf, data[1:])

	pdf = utils.Image(pdf)

	if pdf.Err() {
		log.Fatalf("Problemas creando archivo PDF reporte: %s\n", pdf.Error())
	}

	err := utils.SavePDF(pdf)
	if err != nil {
		log.Fatalf("No se puede guardar archivo PDF: %s|n", err)
	}

}
