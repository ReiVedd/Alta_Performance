package main

import "fmt"

func Depositar(saldo *float64, valor float64) {
	*saldo += valor
}

func Sacar(saldo *float64, valor float64) {
	if valor > *saldo {
		fmt.Println("Saldo insuficiente para saque.")
		return
	}
	*saldo -= valor
}

func main() {

	var saldoDaConta float64 = 100.00

	fmt.Println("Saldo inicial da conta = ", saldoDaConta)
	Depositar(&saldoDaConta, 50.00)

	fmt.Println("Saldo da conta após depósito = ", saldoDaConta)
	Sacar(&saldoDaConta, 30.00)

	fmt.Println("Saldo da conta após saque= ", saldoDaConta)
	Sacar(&saldoDaConta, 200.00)

	fmt.Println("Saldo da conta após saque insuficiente= ", saldoDaConta)
}
