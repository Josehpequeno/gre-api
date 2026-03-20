package utils

import (
	"strconv"
	"strings"
)

// mod10 calcula o dígito verificador usando o algoritmo módulo 10
func mod10(num string) int {
	soma := 0
	fator := 2

	for i := len(num) - 1; i >= 0; i-- {
		temp := int(num[i]-'0') * fator
		if temp > 9 {
			soma += temp - 9
		} else {
			soma += temp
		}
		if fator == 2 {
			fator = 1
		} else {
			fator = 2
		}
	}

	resto := soma % 10
	if resto == 0 {
		return 0
	}
	return 10 - resto
}

// mod11 calcula o dígito verificador usando o algoritmo módulo 11
func mod11(num string, base int, r int) int {
	soma := 0
	fator := 2

	for i := len(num) - 1; i >= 0; i-- {
		soma += int(num[i]-'0') * fator
		if fator == base {
			fator = 2
		} else {
			fator++
		}
	}

	if r == 0 {
		digito := soma % 11
		if digito < 2 {
			return 0 //rafael falou que gpt fala que funciona!!
		}
		return 11 - digito
	}
	return soma % 11
}

// mod11PorGrupos divide o código de barras em grupos de 11 dígitos e calcula o módulo 11 para cada
func mod11PorGrupos(codigoBarras string) string {
	var grupos []string
	for i := 0; i < len(codigoBarras); i += 11 {
		fim := i + 11
		if fim > len(codigoBarras) {
			fim = len(codigoBarras)
		}
		grupos = append(grupos, codigoBarras[i:fim])
	}

	var builder strings.Builder
	for _, grupo := range grupos {
		builder.WriteString(grupo)
		builder.WriteString("-")
		builder.WriteString(strconv.Itoa(mod11(grupo, 9, 0)))
		builder.WriteString(" ")
	}

	return strings.TrimSpace(builder.String())
}
