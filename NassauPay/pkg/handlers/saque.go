package handlers

import (
	"net/http"
	"fmt"
)

type SaqueHandler struct {
	ID string
	ContaOrigem string
	Valor float64
}

func SaqueHandlerFunc(w http.ResponseWriter, r *http.Request) {
	dadosDoSaque := SaqueHandler{
		ID: "TX-54321",
		ContaOrigem: "Conta-Corrente-Nassau",
		Valor: 200.00,
	}
	resposta := fmt.Sprintf(
		"NassauPay Engine:\nSaque ID: %s\nConta de Origem: %s\nValor: R$ %.2f",
		dadosDoSaque.ID,
		dadosDoSaque.ContaOrigem,
		dadosDoSaque.Valor,
	)
	fmt.Fprint(w, resposta)
}