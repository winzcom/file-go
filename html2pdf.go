package html2pdf

import (
	"log"
	"strings"

	wkhtmltopdf "github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

func Html2PDF(html string) []byte {
	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		log.Fatalln("Failed to load library")
	}
	pdfg.AddPage(wkhtmltopdf.NewPageReader(strings.NewReader(html)))

	pdfg.Create()

	return pdfg.Buffer().Bytes()
}
