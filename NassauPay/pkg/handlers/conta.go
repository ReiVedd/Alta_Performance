package handlers

import (
	"fmt"
	"net/http"
	"strconv"
)

type Conta struct {
	Numero      string `json:"numero"`
	NomeTitular string `json:"nome_titular"`
	Saldo       int64  `json:"saldo"`
}

var (
	Contas   = make(map[string]*Conta)// <-- IMPORTANTE: Adicionado de volta para proteger o mapa
)

func ContaHandlerFunc(w http.ResponseWriter, r *http.Request) {
	// Garante que a criação exige POST
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido. Use POST para criar contas.", http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Error(w, "Erro ao ler os dados do formulário", http.StatusBadRequest)
		return
	}

	numero := r.FormValue("numero")
	titular := r.FormValue("titular")
	saldoStr := r.FormValue("saldo")

	if numero == "" || titular == "" {
		http.Error(w, "Número da conta e Nome do Titular são obrigatórios", http.StatusBadRequest)
		return
	}

	saldoInicial, err := strconv.ParseInt(saldoStr, 10, 64)
	if err != nil || saldoInicial < 0 {
		saldoInicial = 0
	}

	novaConta := &Conta{
		Numero:      numero,
		NomeTitular: titular,
		Saldo:       saldoInicial,
	}
	Contas[numero] = novaConta // Destranca

	saldoEmReal := float64(novaConta.Saldo) / 100.0
	
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "NassauPay Engine:\nConta %s criada com sucesso para %s!\nSaldo Inicial: R$ %.2f", 
		novaConta.Numero, novaConta.NomeTitular, saldoEmReal)
}

// ConsultarContaHandler busca uma conta no mapa e mostra os dados na tela
func ConsultarContaHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Método não permitido. Use GET para consultar.", http.StatusMethodNotAllowed)
		return
	}

	numeroConta := r.URL.Query().Get("numero")
	if numeroConta == "" {
		http.Error(w, "O parâmetro 'numero' é obrigatório na URL", http.StatusBadRequest)
		return
	}

// Acesso direto ao mapa, sem tranca nenhuma
conta, existe := Contas[numeroConta] 

if !existe {
    http.Error(w, "Conta não encontrada", http.StatusNotFound)
    return
}

	// 5. Agora sim, o Go sabe quem é 'conta' e consegue pegar o Saldo!
	saldoEmReal := float64(conta.Saldo) / 100.0

	// 6. Printamos os dados reais que estão guardados na memória
	resposta := fmt.Sprintf(
		"--- NassauPay - Consulta de Conta ---\nNúmero: %s\nTitular: %s\nSaldo Atual: R$ %.2f\n-------------------------------------",
		conta.Numero,
		conta.NomeTitular,
		saldoEmReal,
	)

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, resposta)
}