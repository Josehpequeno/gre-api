package utils

import (
	"bytes"
	"gre-api/internal/models"
	"html/template"
	"path/filepath"
	"time"
)

func RenderizarTemplate(req models.BoletoRequest, gre *models.BoletoResponse, pixBase64 string) (string, error) {
	tmpl, err := template.ParseFiles(filepath.Join("internal", "templates", "boleto.html"))
	if err != nil {
		return "", err
	}

	data := struct {
		models.BoletoRequest
		*models.BoletoResponse
		PixBase64      string
		DataGeracao    string
		ValorFormatado string
	}{
		BoletoRequest:  req,
		BoletoResponse: gre,
		PixBase64:      pixBase64,
		DataGeracao:    time.Now().Format("02/01/2006"),
		ValorFormatado: req.Valor,
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return "", err
	}

	return buf.String(), nil
}
