package main

import (
	"fmt"
)

func main() {
	contaDaSilvia := ContaCorrente{titular: "Silvia", saldo: 300}
	contaDoGustavo := ContaCorrente{titular: "Gustavo", saldo: 100}

	fmt.Println(contaDaSilvia)
	fmt.Println(contaDoGustavo)

	status := contaDoGustavo.Transferir(-200, &contaDaSilvia)

	fmt.Println(contaDaSilvia)
	fmt.Println(contaDoGustavo)

	fmt.Println(status)

}

// func main() {

// 	// var contaDoXavier ContaCorrente = ContaCorrente{}

// 	// contaDoXavier := ContaCorrente{titular: "Xavier", saldo: 1000} // posso usar quando não for utilizar todos os campos do objeto, por exemplo
// 	// contaDoXavier := ContaCorrente{"Xavier", 3094, 236613, 1000} // caso utilize todos os campos

// 	// contaDoXavier := ContaCorrente{titular: "Xavier", numeroAgencia: 3094, numeroConta: 236613, saldo: 1000}

// 	// fmt.Println(contaDoXavier == contaDoXavier2)

// 	// // var contaDaCris ContaCorrente
// 	var contaDaCris *ContaCorrente // possibilita a criação da variável pq já reserva um local específico na memória para esse objeto
// 	contaDaCris = new(ContaCorrente)
// 	contaDaCris.titular = "Cris"
// 	contaDaCris.saldo = 500

// 	var contaDaCris2 *ContaCorrente // possibilita a criação da variável pq já reserva um local específico na memória para esse objeto
// 	contaDaCris2 = new(ContaCorrente)
// 	contaDaCris2.titular = "Cris"
// 	contaDaCris2.saldo = 500

// 	fmt.Println(*contaDaCris == *contaDaCris2)
// 	// fmt.Println(*contaDaCris)
// }

// func main() {

// 	contaDaSilvia := ContaCorrente{}
// 	contaDaSilvia.titular = "Silvia"
// 	contaDaSilvia.saldo = 500

// 	fmt.Println(contaDaSilvia)

// 	// valorDoSaque := 200. // . para indicar que é do tipo float
// 	// contaDaSilvia.saldo = contaDaSilvia.saldo - valorDoSaque

// 	fmt.Println(contaDaSilvia.saldo)

// 	// fmt.Println(contaDaSilvia.Sacar(-100))
// 	// fmt.Println(contaDaSilvia.Depositar(1000))
// 	status, valor := contaDaSilvia.Depositar(1000)
// 	fmt.Println(status, valor)

// }
