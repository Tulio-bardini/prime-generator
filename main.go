package main

import (
	"fmt"
	"math/big"
	"time"
)

func main() {

	var n int
	// tamanho do número
	n = 4096

	start2 := time.Now()
	var bigInt *big.Int = generateRandomLaggedNumber(n)
	elapsed2 := time.Now().Sub(start2)
	fmt.Printf("Tempo Lagged Generator: %v \n", elapsed2)

	start2 = time.Now()
	bigInt = generateRandomBlumBlumNumber(n)
	elapsed2 = time.Now().Sub(start2)
	fmt.Printf("Tempo BlumBlum Generator: %v \n", elapsed2)

	start2 = time.Now()
	bigInt = generateRandomLaggedNumber(n)
	elapsed2 = time.Now().Sub(start2)
	fmt.Printf("Tempo Lagged Generator: %v \n", elapsed2)

	start2 = time.Now()
	bigInt = generateRandomBlumBlumNumber(n)
	elapsed2 = time.Now().Sub(start2)
	fmt.Printf("Tempo BlumBlum Generator:: %v \n", elapsed2)

	bigInt = generateRandomBlumBlumNumber(n)
	start1 := time.Now()
	bigInt1 := generatePrimeFermat(n, bigInt)
	elapsed1 := time.Now().Sub(start1)
	fmt.Printf("Tempo Fermat com Blum Blum: %v \n", elapsed1)
	fmt.Printf("Numero Fermat com Blum Blum: %v \n", bigInt1)

	bigInt = generateRandomLaggedNumber(n)
	start2 = time.Now()
	bigInt2 := generatePrimeFermat(n, bigInt)
	elapsed2 = time.Now().Sub(start2)
	fmt.Printf("Tempo Fermat com Lagged: %v \n", elapsed2)
	fmt.Printf("Numero Fermat com Lagged: %v \n", bigInt2)

	bigInt = generateRandomBlumBlumNumber(n)
	start2 = time.Now()
	bigInt2 = generatePrimeMillerRabin(n, bigInt)
	elapsed2 = time.Now().Sub(start2)
	fmt.Printf("Tempo Miller-Rabin com Blum Blum: %v \n", elapsed2)
	fmt.Printf("Numero Miller-Rabin com Blum Blum: %v \n", bigInt2)

	bigInt = generateRandomLaggedNumber(n)
	start1 = time.Now()
	bigInt1 = generatePrimeMillerRabin(n, bigInt)
	elapsed1 = time.Now().Sub(start1)
	fmt.Printf("Tempo Miller-Rabin com Lagged: %v \n", elapsed1)
	fmt.Printf("Numero Miller-Rabin com Lagged: %v \n", bigInt1)

}
