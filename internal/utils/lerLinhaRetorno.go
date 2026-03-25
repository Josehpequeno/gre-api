package utils

import (
	"fmt"
	"gre-api/internal/models"
	"log"
	"math"
	"strconv"
	"strings"
	"time"
)

type InformacoesCodigoBarras struct {
	Digito     string
	Valor      string
	Convenio   string
	Vencimento string
	Codigo     string
	Parcela    string
}

// Função para extrair informações do código de barras
func ExtrairInformacoes(codigoBarras string) (*InformacoesCodigoBarras, error) {
	if len(codigoBarras) != 44 {
		return nil, fmt.Errorf("código de barras inválido")
	}
	// 89900000002059000010109520724010301300659915
	digito := codigoBarras[3:4]
	vlr := codigoBarras[4:15]
	convenio := codigoBarras[22:27]
	vencimento := codigoBarras[27:35]
	codigo := codigoBarras[35:42]
	parcela := codigoBarras[42:44]

	return &InformacoesCodigoBarras{
		Digito:     digito,
		Valor:      vlr,
		Convenio:   convenio,
		Vencimento: vencimento,
		Codigo:     codigo,
		Parcela:    parcela,
	}, nil
}

func LerLinhaRetorno(linha string) (models.LinhaRetorno, error) {
	// --- Helpers internos ---
	parseDate := func(raw string) (string, error) {
		t, err := time.Parse("20060102", raw)
		if err != nil {
			return "", fmt.Errorf("erro ao converter data %q: %w", raw, err)
		}
		return t.Format("02/01/2006"), nil
	}

	parseFloat := func(raw string) (float64, error) {
		val, err := strconv.ParseFloat(strings.TrimSpace(raw), 64)
		if err != nil {
			return 0, fmt.Errorf("erro ao converter número %q: %w", raw, err)
		}
		// sempre duas casas
		return math.Round(val) / 100, nil
	}

	// --- Extrair campos fixos ---
	codigoRegistro := linha[0:1]
	identificacao := linha[1:21]
	dataPagamentoRaw := linha[21:29]
	dataCreditoRaw := linha[29:37]
	codigoDeBarras := linha[37:81]
	valorRaw := linha[81:93]
	tarifaRaw := linha[93:100]
	nsr := linha[100:108]
	agenciaArrecadadora := linha[108:116]
	formaArrecadacao := linha[116:117]
	numeroAutenticacao := linha[117:140]
	formaPagamento := linha[140:141]

	// --- Parse datas ---
	dataPagamento, _ := parseDate(dataPagamentoRaw)
	dataCredito, _ := parseDate(dataCreditoRaw)

	// --- Info código de barras ---
	info, err := ExtrairInformacoes(codigoDeBarras)
	if err != nil {
		log.Println("Erro ao extrair informações do código de barras:", err)
	}

	// --- Identificação de aluno/inscrito ---
	parcelaInt, _ := strconv.Atoi(info.Parcela)

	// --- Valores ---
	valorPago, _ := parseFloat(valorRaw)
	valorBoleto, _ := parseFloat(info.Valor)
	tarifa, _ := parseFloat(tarifaRaw)

	// --- Monta struct final ---
	return models.LinhaRetorno{
		CodigoRegistro:      codigoRegistro,
		Codigo:              info.Codigo,
		ValorBoleto:         valorBoleto,
		ValorPago:           valorPago,
		DataPagamento:       dataPagamento,
		AgenciaArrecadadora: agenciaArrecadadora,
		CodigoBarras:        codigoDeBarras,
		FormaPagamento:      formaPagamento,
		Parcela:             parcelaInt,
		DataCredito:         dataCredito,
		Tarifa:              tarifa,
		NSR:                 nsr,
		FormaArrecadacao:    formaArrecadacao,
		NumeroAutenticacao:  numeroAutenticacao,
		Identificacao:       identificacao,
	}, nil
}

func SplitPorTamanho(s string, tamanho int) []string {
	var linhas []string
	for i := 0; i < len(s); i += tamanho {
		fim := i + tamanho
		if fim > len(s) {
			fim = len(s)
		}
		linhas = append(linhas, s[i:fim])
	}
	return linhas
}
