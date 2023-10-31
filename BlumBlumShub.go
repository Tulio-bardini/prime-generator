package main

import (
	"fmt"
	"math/big"

	// Biblioteca usada para obter a posição do ponteiro do mouse
	"github.com/go-vgo/robotgo"
)

func generateRandomBlumBlumNumber(sizeNumberBits int) *big.Int {

	loopNumber := sizeNumberBits / 8

	// Define as variáveis do algoritmo
	p := 30000000091
	q := 40000000003
	m := big.NewInt(int64(p * q))

	xPosition, yPosition := robotgo.GetMousePos()
	seed := big.NewInt(int64((xPosition + 1471) * (yPosition + 1249)))
	var byteNumber byte
	var numberHexaString string

	// Concatena bit a bit no número pseudo aleatório
	for i := 0; i < loopNumber; i++ {

		// Preenche um byte
		for j := 0; j < 8; j++ {
			seed := seed.Exp(seed, big.NewInt(2), m)
			byteNumber = (byteNumber << 1) | (seed.Bytes()[len(seed.Bytes())-1] & 1)
		}

		// Cria um string com esses bytes hexadecimais gerados
		numberHexaString = numberHexaString + fmt.Sprintf("%02X", byteNumber)
	}

	bigInt := new(big.Int)
	bigInt.SetString(numberHexaString, 16)

	return bigInt
}
