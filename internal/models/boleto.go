package models

type BoletoRequest struct {
	CodigoBanco            string `json:"codigo_banco" example:"001-9" binding:"required"`
	Convenio               string `json:"convenio" example:"12345" binding:"required"`
	BeneficiarioNome       string `json:"beneficiario_nome" example:"Empresa Ltda" binding:"required"`
	BeneficiarioCnpjCpf    string `json:"beneficiario_cnpj_cpf" example:"12.345.678/0001-90" binding:"required"`
	BeneficiarioConta      string `json:"beneficiario_conta" binding:"required"`
	BeneficiarioAgencia    string `json:"beneficiario_agencia" binding:"required"`
	BeneficiarioEndereco   string `json:"beneficiario_endereco" binding:"omitempty"`
	PagadorNome            string `json:"pagador_nome" example:"João da Silva" binding:"required"`
	PagadorCpf             string `json:"pagador_cpf" example:"123.456.789-00" binding:"required"`
	PagadorEndereco        string `json:"pagador_endereco" binding:"omitempty"`
	IdentificadorPagador   int    `json:"identificador_pagador" binding:"omitempty"`
	Valor                  string `json:"valor" example:"150,75" binding:"required"`
	DataVencimento         string `json:"data_vencimento" example:"31/12/2025" binding:"required"`
	NumeroDocumento        string `json:"numero_documento" example:"123456" binding:"required"`
	NossoNumero            string `json:"nosso_numero" example:"123456" binding:"required"`
	InformacaoBeneficiario string `json:"informacao_beneficiario" example:"Multa de 2% após vencimento" binding:"required"`
	Parcela                int    `json:"parcela" example:"1" binding:"required"`
	RetornarPDF            bool   `json:"retornar_pdf" example:"true" `
}

type BoletoResponse struct {
	CodigoBarras         string  `json:"codigoBarras"`
	ImageCodigoDeBarras  string  `json:"imageCodigoDeBarras"`
	LinhaDigitavel       string  `json:"linhaDigitavel"`
	NumeroDocumento      string  `json:"numeroDocumento"`
	NomeBeneficiario     string  `json:"nomeBeneficiario"`
	CpfCnpjBeneficiario  string  `json:"cpfCnpjBeneficiario"`
	EnderecoBeneficiario string  `json:"enderecoBeneficiario"`
	NomeDevedor          string  `json:"nomeDevedor"`
	CpfDevedor           string  `json:"cpfDevedor"`
	EnderecoDevedor      string  `json:"enderecoDevedor"`
	Convenio             string  `json:"convenio"`
	Matricula            string  `json:"matricula"`
	Vencimento           string  `json:"vencimento"`
	Curso                string  `json:"curso"`
	Valor                float64 `json:"valor"`
	PixBase64            string  `json:"pixBase64"`
	HTML                 string  `json:"html,omitempty"`
	PDF                  string  `json:"pdf,omitempty"`
}

type ReadRetornoRequest struct {
	Retorno string `json:"retorno" binding:"required"`
}
type LinhaRetorno struct {
	Codigo                       string  `json:"codigo"`
	CursoID                      int     `json:"curso_id"`
	ValorBoleto                  float64 `json:"valor_boleto"`
	ValorPago                    float64 `json:"valor_pago"`
	DataPagamento                string  `json:"data_pagamento"`
	AgenciaArrecadadora          string  `json:"agencia_arrecadadora"`
	CodigoBarras                 string  `json:"codigo_barras"`
	FormaPagamento               string  `json:"forma_pagamento"`
	Parcela                      int     `json:"parcela"`
	DataCredito                  string  `json:"data_credito"`
	Tarifa                       float64 `json:"tarifa"`
	NSR                          string  `json:"nsr"`
	FormaArrecadacao             string  `json:"forma_arrecadacao"`
	NumeroAutenticacao           string  `json:"numero_autenticacao"`
	MatriculaInscricaoPorExtenso string  `json:"matricula_inscricao_por_extenso"`
}
