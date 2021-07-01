package main

import (
	"fmt"
	"go-101/contas"
)

// Aula 5
func main() {
	contaDoDenis := contas.ContaPoupanca{}
	contaDoDenis.Depositar(100)
	PagarBoleto(&contaDoDenis, 60)

	fmt.Println(contaDoDenis.ObterSaldo())

	contaDaLuisa := contas.ContaCorrente{}
	contaDaLuisa.Depositar(500)
	PagarBoleto(&contaDaLuisa, 400)

	fmt.Println(contaDaLuisa.ObterSaldo())
}

func PagarBoleto(conta verificarConta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}

type verificarConta interface {
	Sacar(valor float64) string
}

// // Aula 4
// func main() {
// 	// clienteBruno := clientes.Titular{"Bruno", "111.222.333.44", "Desenvolvedor"}
// 	// contaDoBruno := contas.ContaCorrente{clienteBruno, 1234, 12345, 1000}

// 	contaExemplo := contas.ContaCorrente{}
// 	contaExemplo.Depositar(100)

// 	fmt.Println(contaExemplo.ObterSaldo())
// }

// Aula 3
// func main() {

// 	contaDaSilvia := contas.ContaCorrente{Titular: "Silvia", Saldo: 300}
// 	contaDoGustavo := contas.ContaCorrente{Titular: "Gustavo", Saldo: 100}

// 	fmt.Println(contaDaSilvia)
// 	fmt.Println(contaDoGustavo)

// 	status := contaDoGustavo.Transferir(-200, &contaDaSilvia)

// 	fmt.Println(contaDaSilvia)
// 	fmt.Println(contaDoGustavo)

// 	fmt.Println(status)
// }

// Aula 2
// func main() {

// 	// var contaDoXavier ContaCorrente = ContaCorrente{}

// 	// contaDoXavier := ContaCorrente{Titular: "Xavier", Saldo: 1000} // posso usar quando não for utilizar todos os campos do objeto, por exemplo
// 	// contaDoXavier := ContaCorrente{"Xavier", 3094, 236613, 1000} // caso utilize todos os campos

// 	// contaDoXavier := ContaCorrente{Titular: "Xavier", numeroAgencia: 3094, numeroConta: 236613, Saldo: 1000}

// 	// fmt.Println(contaDoXavier == contaDoXavier2)

// 	// // var contaDaCris ContaCorrente
// 	var contaDaCris *ContaCorrente // possibilita a criação da variável pq já reserva um local específico na memória para esse objeto
// 	contaDaCris = new(ContaCorrente)
// 	contaDaCris.Titular = "Cris"
// 	contaDaCris.Saldo = 500

// 	var contaDaCris2 *ContaCorrente // possibilita a criação da variável pq já reserva um local específico na memória para esse objeto
// 	contaDaCris2 = new(ContaCorrente)
// 	contaDaCris2.Titular = "Cris"
// 	contaDaCris2.Saldo = 500

// 	fmt.Println(*contaDaCris == *contaDaCris2)
// 	// fmt.Println(*contaDaCris)
// }

// Aula 1
// func main() {

// 	contaDaSilvia := ContaCorrente{}
// 	contaDaSilvia.Titular = "Silvia"
// 	contaDaSilvia.Saldo = 500

// 	fmt.Println(contaDaSilvia)

// 	// valorDoSaque := 200. // . para indicar que é do tipo float
// 	// contaDaSilvia.Saldo = contaDaSilvia.Saldo - valorDoSaque

// 	fmt.Println(contaDaSilvia.Saldo)

// 	// fmt.Println(contaDaSilvia.Sacar(-100))
// 	// fmt.Println(contaDaSilvia.Depositar(1000))
// 	status, valor := contaDaSilvia.Depositar(1000)
// 	fmt.Println(status, valor)
// }
