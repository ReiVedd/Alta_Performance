package handlers

import (
	"net/http"
	"fmt"
)

type DepositoHandler struct {
	ID string
	ContaDestino string
	Valor float64
}

func DepositoHandlerFunc(w http.ResponseWriter, r *http.Request) {
	dadosDoDeposito := DepositoHandler{
		ID: "TX-12345",
		ContaDestino: "Conta-Corrente-Nassau",
		Valor: 500.004897665,
	}

		resposta := fmt.Sprintf(
		"NassauPay Engine:\nDepósito ID: %s\nDestino: %s\nValor: R$ %.2f ",
		dadosDoDeposito.ID,
		dadosDoDeposito.ContaDestino,
		dadosDoDeposito.Valor,
	)
	fmt.Fprint(w, resposta)
}
