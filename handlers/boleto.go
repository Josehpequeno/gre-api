package handlers

import (
	"bytes"
	"gre-api/models"
	"html/template"
	"net/http"
	"path/filepath"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"github.com/gin-gonic/gin"
)

// RenderBoleto gera o boleto em HTML a partir dos dados recebidos.
// @Summary Gera um boleto bancário
// @Description Recebe os dados do boleto e retorna o HTML gerado
// @Tags boletos
// @Accept json
// @Produce html
// @Param request body models.BoletoRequest true "Dados do boleto"
// @Success 200 {object} BoletoResponse
// @Failure 400 {object} map[string]string "Erro de validação"
// @Failure 500 {object} map[string]string "Erro interno"
// @Router /boleto [post]
func RenderBoleto(c *gin.Context) {
	var req models.BoletoRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tmplPath := filepath.Join("internal", "templates", "boleto.html")
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao carregar template: " + err.Error()})
		return
	}

	var htmlBuf bytes.Buffer
	if err := tmpl.Execute(&htmlBuf, req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao renderizar HTML: " + err.Error()})
		return
	}

	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar gerador PDF" + err.Error()})
		return
	}

	// Configurações opcionais
	pdfg.Dpi.Set(300)
	pdfg.Orientation.Set(wkhtmltopdf.OrientationPortrait)
	pdfg.Grayscale.Set(false)

	// Adicionar página com o HTML
	page := wkhtmltopdf.NewPageReader(bytes.NewReader(htmlBuf.Bytes()))
	// Se houver assets (CSS, imagens), defina o caminho base
	// page.FooterRight.Set("[page]")
	pdfg.AddPage(page)

	// Gerar PDF
	if err := pdfg.Create(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao gerar PDF: " + err.Error()})
		return
	}

	// 3. Retornar o PDF como resposta
	c.Header("Content-Disposition", "attachment; filename=boleto.pdf")
	c.Data(http.StatusOK, "application/pdf", pdfg.Bytes())
}
