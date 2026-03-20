package utils

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"gre-api/internal/models"
	"image/png"
	"log"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/code128"
)

type gerarCodigoParms struct {
	identificacao int
	convenio      int
	valor         float64
	vencimento    string
	parcela       int
}

type InformaçoesCodigoBarras struct {
	Digito     string
	Valor      string
	Convenio   string
	Vencimento string
	Codigo     string
	Parcela    string
}

func zeroEsquerda(numero int, comprimento int) string {
	return fmt.Sprintf("%0*d", comprimento, numero)
}

// gerarCodigo gera o código de barras e linha digitável
func gerarCodigo(g gerarCodigoParms) (map[string]string, error) {
	var vlr string

	vlr = fmt.Sprintf("%011d", int(g.valor*100))

	var codigo string
	parcelaStr := zeroEsquerda(g.parcela, 2)
	codigo = fmt.Sprintf("%s%s", zeroEsquerda(g.identificacao, 7), parcelaStr)

	base := fmt.Sprintf("899%s0001010%d%s%s", vlr, g.convenio, g.vencimento, codigo)
	// base := fmt.Sprintf("899%s0001010%d20250510%s", vlr, convenio, codigo)
	digito := mod11(base, 9, 0)
	codigoBarras := fmt.Sprintf("899%d%s0001010%d%s%s", digito, vlr, g.convenio, g.vencimento, codigo)
	linhadigitavel := mod11PorGrupos(codigoBarras)
	numeroDocumento := gerarNumeroDocumento(g.parcela)

	return map[string]string{
		"LinhaDigitavel":  linhadigitavel,
		"CodigoBarras":    codigoBarras,
		"NumeroDocumento": numeroDocumento,
	}, nil
}

func gerarNumeroDocumento(parcela int) string {
	// Data atual no formato YYYYMMDD
	dataAtual := time.Now().Format("20060102")

	// Número aleatório de 6 dígitos
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	numeroAleatorio := r.Intn(900000) + 100000

	// Formata a parcela com 2 dígitos (ex: 01, 02, ..., 99)
	numeroParcela := fmt.Sprintf("%02d", parcela)

	// Combina data, parcela e número aleatório
	return fmt.Sprintf("%s%s%d", dataAtual, numeroParcela, numeroAleatorio)
}

// gerarImagemCodigoDeBarras gera a imagem do código de barras em base64
func gerarImagemCodigoDeBarras(codigo string) (string, error) {
	// Cria o código de barras
	bc, err := code128.Encode(codigo)
	if err != nil {
		return "", fmt.Errorf("erro ao codificar código de barras: %v", err)
	}

	// Redimensiona (opcional)
	scaledBC, err := barcode.Scale(bc, bc.Bounds().Dx(), 100)
	if err != nil {
		return "", fmt.Errorf("erro ao redimensionar código de barras: %v", err)
	}
	// Use scaledBC directly without reassigning to bc
	if scaledBC == nil {
		return "", fmt.Errorf("erro ao redimensionar código de barras: código de barras redimensionado é nulo")
	}

	// Codifica para PNG
	var buf bytes.Buffer
	err = png.Encode(&buf, bc)
	if err != nil {
		return "", fmt.Errorf("erro ao codificar PNG: %v", err)
	}

	// Converte para base64
	return base64.StdEncoding.EncodeToString(buf.Bytes()), nil
}

// GerarGRE é a função principal que orquestra a geração do boleto
func GerarGRE(boletoReq models.BoletoRequest) (map[string]string, error) {
	convenioInt, err := strconv.Atoi(boletoReq.Convenio)
	if err != nil {
		return nil, fmt.Errorf("convenio inválido: %w", err)
	}

	vencimentoBoleto, err := FormatarData(boletoReq.DataVencimento)
	if err != nil {
		return nil, fmt.Errorf("data de vencimento inválida: %w", err)
	}

	valorStringPonto := strings.ReplaceAll(boletoReq.Valor, ",", ".")
	parsedValor, err := strconv.ParseFloat(valorStringPonto, 64)
	if err != nil {
		return nil, fmt.Errorf("erro ao converter valor para float: %w", err)
	}

	gerarCodigoParms := gerarCodigoParms{
		identificacao: boletoReq.IdentificadorPagador,
		convenio:      convenioInt,
		valor:         parsedValor,
		vencimento:    vencimentoBoleto,
		parcela:       boletoReq.Parcela,
	}
	resultado, err := gerarCodigo(gerarCodigoParms)
	if err != nil {
		log.Fatal("erro ao gerar código:", err)
		return nil, err
	}

	imagemCodigo, err := gerarImagemCodigoDeBarras(resultado["CodigoBarras"])
	if err != nil {
		log.Fatal("erro ao gerar imagem código de barras:", err)
		return nil, err
	}

	// return nil, errors.New("erro ao gerar imagem código de barras" + imagemCodigo)

	resultado["ImageCodigoDeBarras"] = imagemCodigo
	return resultado, nil
}
