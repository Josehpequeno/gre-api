package handlers

import (
	"gre-api/internal/models"
	"gre-api/internal/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Gera um boleto bancário
// @Description Retorna HTML ou PDF dependendo do parâmetro RetornarPDF
// @Tags boletos
// @Accept json
// @Produce text/html
// @Produce application/pdf
// @Param request body models.BoletoRequest true "Dados do boleto"
// @Success 200 {string} string "HTML do boleto"
// @Success 200 {file} file "PDF do boleto"
// @Failure 400 {object} map[string]string "Erro de validação"
// @Failure 500 {object} map[string]string "Erro interno"
// @Router /boleto [post]
func RenderBoleto(c *gin.Context) {
	var req models.BoletoRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	greResponse, err := utils.GerarGRE(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao gerar GRE"})
		return
	}
	// Simula a geração do PIX (substitua pela implementação real)
	pixBase64, err := utils.CriarPix(req)
	if err != nil {
		log.Printf("Erro ao gerar PIX: %v\n", err)
		pixBase64 = ""
	}

	// Gera o HTML do boleto
	gre := &models.BoletoResponse{
		CodigoBarras:        greResponse["CodigoBarras"],
		ImageCodigoDeBarras: greResponse["ImageCodigoDeBarras"],
		LinhaDigitavel:      greResponse["LinhaDigitavel"],
		NumeroDocumento:     greResponse["NumeroDocumento"],
	}
	html, err := utils.RenderizarTemplate(req, gre, pixBase64)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao renderizar boleto"})
		return
	}

	if req.RetornarPDF {
		pdf, err := utils.GerarPDF(html, req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao gerar PDF"})
			return
		}
		// Define os cabeçalhos para download de arquivo PDF
		c.Header("Content-Disposition", "attachment; filename=boleto.pdf")
		c.Header("Content-Type", "application/pdf")
		// Envia o conteúdo do PDF como resposta
		c.Data(http.StatusOK, "application/pdf", []byte(pdf))
	} else {
		c.Header("Content-Type", "text/html; charset=utf-8")
		c.String(http.StatusOK, html)
	}
}

// ReadRetorno godoc
// @Summary      Processa arquivo de retorno
// @Description  Lê o conteúdo do arquivo de retorno bancário e extrai os boletos
// @Tags         retorno
// @Accept multipart/form-data
// @Produce      json
// @Param retorno formData string true "Conteúdo do arquivo retorno (texto com múltiplas linhas)"
// @Success      200 {array} models.LinhaRetorno
// @Failure      400 {object} map[string]string
// @Router       /retorno/read [post]
func ReadRetorno(c *gin.Context) {
	retorno := c.PostForm("retorno")
	resultados := make([]models.LinhaRetorno, 0)
	it := 1
	linhas := utils.SplitPorTamanho(retorno, 150)
	log.Println("linhas", linhas)
	for _, linha := range linhas {
		if linha == "" {
			continue
		}
		log.Println("linha", linha)

		if it == 2 {
			if len(linha) >= 150 {
				infoLinha, err := utils.LerLinhaRetorno(linha)
				if err != nil {
					log.Println("Erro ao ler linha de retorno:", err)
					continue
				}
				log.Println(infoLinha)
				resultados = append(resultados, infoLinha)
			} else {
				log.Println("Linha ignorada (não há palavra suficiente):", linha)
			}
		}
		if it%3 == 0 {
			log.Printf("fim de boleto %d", it)
			log.Println()
			it = 1
		}
		it++
	}
	log.Println("resultados", resultados)
	c.JSON(http.StatusOK, resultados)
}
