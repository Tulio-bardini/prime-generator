package main

import (
	"crypto/rand"
	"math"
	"math/big"
)

func millerRabin(number *big.Int, k int) bool {

	// Checa se o número é par
	if number.Bytes()[len(number.Bytes())-1]&1 == 0 {
		return false
	}

	r := 1.0
	// Soma -1
	oneLess := new(big.Int)
	oneLess.SetInt64(-1)

	d := new(big.Int)
	d.Add(number, oneLess)
	result := new(big.Int)

	for { // Encontra o número na forma d*2^r, onde d é impar
		expo := big.NewInt(int64(math.Pow(2, r)))
		result.Mod(d, expo)
		// Checa se o resto da divisão é zero
		if result.Cmp(big.NewInt(0)) == 0 {
			result.Div(d, expo)
			// Checa se o numero é impar
			if (result.Bytes()[len(result.Bytes())-1] & 1) == 1 {
				d.Div(d, expo)
				break
			}
		}
		r += 1
	}

	millerTest := func(r int) bool {
		// gera um número aleatório
		checkPrimality, _ := rand.Int(rand.Reader, number)
		checkPrimality.Add(checkPrimality, oneLess)

		// (checkPrimality**d) mod number
		checkPrimality.Exp(checkPrimality, d, number)

		// Verifica se checkPrimality é igual a 1
		if checkPrimality.Cmp(big.NewInt(1)) == 0 {
			return true
		}

		for count := 0; count < int(r); {
			// Verifica se checkPrimality é igual a n-1
			if checkPrimality.Cmp(new(big.Int).Add(number, oneLess)) == 0 {
				return true
			} else {
				checkPrimality.Exp(checkPrimality, big.NewInt(2), number)
			}
			count += 1
		}
		return false
	}

	for i := 0; i < k; i++ { // Roda o millerRabin k vezes
		if millerTest(int(r)) == false {
			return false
		}
	}

	return true
}

func fermatTest(number *big.Int) bool {

	a := new(big.Int)
	oneLess := new(big.Int)
	oneLess.SetInt64(-1)

	// Encontra um numero que é primo entre si com o number
	a, _ = rand.Int(rand.Reader, number)
	test := new(big.Int).GCD(nil, nil, a, number) // MDC

	if test.Cmp(big.NewInt(1)) != 0 {
		return false
	}

	// Calcula a^(p-1) mod p
	a.Exp(a, new(big.Int).Add(number, oneLess), number)

	if a.Cmp(big.NewInt(1)) == 0 {
		return true
	}

	return false
}
