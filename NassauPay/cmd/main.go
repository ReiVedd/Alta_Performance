package main

import (
	"net/http"
	"fmt"
	"NassauPay/pkg/handlers"
)

func main() {
	http.HandleFunc("/deposito", handlers.DepositoHandlerFunc)
	http.HandleFunc("/conta", handlers.ContaHandlerFunc)
	http.HandleFunc("/saque", handlers.SaqueHandlerFunc)

	http.HandleFunc("/consultar", handlers.ConsultarContaHandler)
	// Invoke-RestMethod -Method Post -Uri http://localhost:8080/conta -Body "numero= &titular=  &saldo= 
	//Pra poder criar a conta via cmd
	//http://localhost:8080/consultar?numero= link para poder verificar a conta criada, so alterar o numero da conta criada

	fmt.Println("NassauPay Engine rodando na porta :8080...")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("Erro ao iniciar o servidor: %v\n", err)
	}
}