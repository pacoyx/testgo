package utils

import (
	"encoding/csv"
	"log"
	"os"
	"time"

	"github.com/jung-kurt/gofpdf"
)

//https://appliedgo.net/pdf/
//https://github.com/jung-kurt/gofpdf/blob/master/fpdf_test.go

//LoadCSV carga csv
func LoadCSV(path string) [][]string {
	f, err := os.Open(path)
	if err != nil {
		log.Fatalf("Cannot open '%s': %s\n", path, err.Error())
	}
	defer f.Close()
	r := csv.NewReader(f)
	rows, err := r.ReadAll()
	if err != nil {
		log.Fatalln("Cannot read CSV data:", err.Error())
	}
	return rows
}

//Path devuelve la ruta del archivo csv
func Path() string {
	if len(os.Args) < 2 {
		return `E:\proyectos.csv`
	}
	return os.Args[1]
}

//NewReport crea nuevo archivo pdf en memoria
func NewReport() *gofpdf.Fpdf {
	pdf := gofpdf.New("L", "mm", "Letter", "")
	pdf.AddPage()
	pdf.SetFont("Times", "B", 28)
	pdf.Cell(40, 10, "Daily Report")
	pdf.Ln(12)
	pdf.SetFont("Times", "", 20)
	pdf.Cell(40, 10, time.Now().Format("Mon Jan 2, 2006"))
	pdf.Ln(20)
	return pdf
}

//Header cabecera
func Header(pdf *gofpdf.Fpdf, hdr []string) *gofpdf.Fpdf {
	pdf.SetFont("Times", "B", 16)
	pdf.SetFillColor(240, 240, 240)
	for _, str := range hdr {
		pdf.CellFormat(40, 7, str, "1", 0, "", true, 0, "")
	}
	pdf.Ln(-1)
	return pdf
}

//Table seccion tabla
func Table(pdf *gofpdf.Fpdf, tbl [][]string) *gofpdf.Fpdf {
	pdf.SetFont("Times", "", 16)
	pdf.SetFillColor(255, 255, 255)
	align := []string{"L", "C", "L", "R", "R", "R"}
	for _, line := range tbl {
		for i, str := range line {
			pdf.CellFormat(40, 7, str, "1", 0, align[i], false, 0, "")
		}
		pdf.Ln(-1)
	}
	return pdf
}

//Image para insert imagen en archivo
func Image(pdf *gofpdf.Fpdf) *gofpdf.Fpdf {
	pdf.ImageOptions("golang.jpg", 225, 10, 25, 25, false, gofpdf.ImageOptions{ImageType: "JPG", ReadDpi: true}, 0, "")
	return pdf
}

//SavePDF guarda archivo pdf en disco
func SavePDF(pdf *gofpdf.Fpdf) error {
	return pdf.OutputFileAndClose("pdfs/" + "report.pdf")
}

//CrearDirectorioSiNoExiste crea la carpeta en la ruta si no existe
func CrearDirectorioSiNoExiste(directorio string) {
	if _, err := os.Stat(directorio); os.IsNotExist(err) {
		err = os.Mkdir(directorio, 0755)
		if err != nil {
			// Aquí puedes manejar mejor el error, es un ejemplo
			panic(err)
		}
	}
}
