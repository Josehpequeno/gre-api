package models

type BoletoRequest struct {
	BeneficiarioNome       string `json:"beneficiario_nome" example:"Empresa Ltda" binding:"required"`
	BeneficiarioCnpjCpf    string `json:"beneficiario_cnpj_cpf" example:"12.345.678/0001-90" binding:"required"`
	BeneficiarioConta      string `json:"beneficiario_conta" binding:"required"`
	BeneficiarioAgencia    string `json:"beneficiario_agencia" binding:"required"`
	PagadorNome            string `json:"pagador_nome" example:"João da Silva" binding:"required"`
	PagadorCpf             string `json:"pagador_cpf" example:"123.456.789-00" binding:"required"`
	PagadorEndereco        string `json:"pagador_endereco" binding:"required"`
	IdentificadorPagador   string `json:"identificador_pagador" binding:"omitempty"`
	Valor                  string `json:"valor" example:"150,75" binding:"required"`
	DataVencimento         string `json:"data_vencimento" example:"31/12/2025" binding:"required"`
	NumeroDocumento        string `json:"numero_documento" example:"123456" binding:"required"`
	NossoNumero            string `json:"nosso_numero" example:"123456" binding:"required"`
	InformacaoBeneficiario string `json:"informacao_beneficiario" example:"Multa de 2% após vencimento" binding:"required"`
}

type BoletoResponse struct {
	HTML string `json:"html"`
	PDF  string `json:"pdf_base64"` // PDF codificado em base64
}
