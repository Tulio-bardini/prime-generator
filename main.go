package main

import (
	"fmt"
	"math/big"
	"time"
)

func main() {

	var n int
	n = 4096

	start2 := time.Now()
	var bigInt *big.Int = generateRandomLaggedNumber(n)
	elapsed2 := time.Now().Sub(start2)
	fmt.Println(elapsed2)

	// start2 := time.Now()
	// bigInt2 := generatePrimeMillerRabin(n, bigInt)
	// elapsed2 := time.Now().Sub(start2)
	// fmt.Println(bigInt2)
	// fmt.Println(elapsed2)

	start2 = time.Now()
	bigInt = generateRandomBlumBlumNumber(n)
	elapsed2 = time.Now().Sub(start2)
	fmt.Println(elapsed2)

	start2 = time.Now()
	bigInt = generateRandomLaggedNumber(n)
	elapsed2 = time.Now().Sub(start2)
	fmt.Println(elapsed2)

	// start1 := time.Now()
	// bigInt1 := generatePrimeFermat(n, bigInt)
	// elapsed1 := time.Now().Sub(start1)
	// fmt.Println(bigInt1)
	// fmt.Println(elapsed1)

	// start2 = time.Now()
	// bigInt2 = generatePrimeFermat(n, bigInt)
	// elapsed2 = time.Now().Sub(start2)
	// fmt.Println(bigInt2)
	// fmt.Println(elapsed2)

	start2 = time.Now()
	bigInt = generateRandomBlumBlumNumber(n)
	elapsed2 = time.Now().Sub(start2)
	fmt.Println(elapsed2)

	// start1 = time.Now()
	// bigInt1 = generatePrimeMillerRabin(n, bigInt)
	// elapsed1 = time.Now().Sub(start1)
	// fmt.Println(bigInt1)
	// fmt.Println(elapsed1)

	fmt.Println(bigInt)
}
