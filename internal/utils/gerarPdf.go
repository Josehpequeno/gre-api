package utils

import (
	"gre-api/internal/models"
	"html/template"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/proto"
)

var baseDir string

func init() {
	// Obter o diretório atual
	var err error
	baseDir, err = os.Getwd()
	if err != nil {
		log.Fatal("Erro ao obter diretório atual:", err)
	}
}

// Método alternativo para carregar templates sem embed
func _(file string) (*template.Template, error) {
	templatePath := filepath.Join(baseDir, "templates", file)
	return template.ParseFiles(templatePath)
}

func GerarPDF(html string, data models.BoletoRequest) (string, error) {

	browser := rod.New().MustConnect()
	defer browser.MustClose()
	page := browser.MustPage("")

	defer page.MustClose()

	page.MustEval(`(html) => {
		document.open()
		document.write(html)
		document.close()
	}`, html)

	pdfOpts := &proto.PagePrintToPDF{
		Landscape:       false,
		PrintBackground: true,
		Scale:           ptrFloat(1.0),
		PaperWidth:      ptrFloat(8.5),
		PaperHeight:     ptrFloat(11),
		MarginTop:       ptrFloat(0.0),
		MarginBottom:    ptrFloat(0.0),
		MarginLeft:      ptrFloat(0.0),
		MarginRight:     ptrFloat(0.0),
	}

	pdfData, err := page.PDF(pdfOpts)
	if err != nil {
		return "", err
	}

	pdfBytes, err := io.ReadAll(pdfData)
	if err != nil {
		return "", err
	}

	return string(pdfBytes), nil
}
