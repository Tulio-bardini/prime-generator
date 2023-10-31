package main

import (
	"math/big"
)

func generatePrimeFermat(bitSize int, bigInt *big.Int) *big.Int {

	// var bigInt *big.Int = generateRandomLaggedNumber(bitSize)
	// var bigInt *big.Int = generateRandomBlumBlumNumber(bitSize)

	if bigInt.Bytes()[len(bigInt.Bytes())-1]&1 == 0 {
		bigInt.Add(bigInt, big.NewInt(1))
	}

	// Utiliza o fermat como teste rapido
	// Executa o miller rabin em 8 iterações, para aumentar a certeza sobre número primo
	for {
		if fermatTest(bigInt) {
			return bigInt
		}
		bigInt.Add(bigInt, big.NewInt(2))
	}

}

func generatePrimeMillerRabin(bitSize int, bigInt *big.Int) *big.Int {

	// var bigInt *big.Int = generateRandomLaggedNumber(bitSize)
	// var bigInt *big.Int = generateRandomBlumBlumNumber(bitSize)

	if bigInt.Bytes()[len(bigInt.Bytes())-1]&1 == 0 {
		bigInt.Add(bigInt, big.NewInt(1))
	}

	// Utiliza o fermat como teste rapido
	// Executa o miller rabin em 8 iterações, para aumentar a certeza sobre número primo
	for {
		if millerRabin(bigInt, 8) {
			return bigInt
		}
		bigInt.Add(bigInt, big.NewInt(2))
	}

}
