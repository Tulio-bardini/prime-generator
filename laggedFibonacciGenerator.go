package main

import (
	"fmt"
	"math/big"
	"strconv"

	// Biblioteca usada para obter a posição do ponteiro do mouse
	"github.com/go-vgo/robotgo"
)

func generateRandomLaggedNumber(sizeNumberBits int) *big.Int {

	loopNumber := sizeNumberBits / 8

	// Variáveis do algoritmo
	j := 2
	k := 6
	m := 16
	xPosition, yPosition := robotgo.GetMousePos()
	xPosition += 9811
	yPosition += 14177
	seed := xPosition * yPosition

	// Tranforma o numero gerado em um numero hexadecimal pelo menos 7 casas
	strNumber := strconv.FormatInt(int64(seed), 16)
	var numberHexaString string

	// Mascaras para pegar os 4 ultimos bits de um byte aleatório gerado
	maskLastBits := byte(0x0F)

	// Cada loop gera 8 bits aleatórios
	for i := 0; i < loopNumber; i++ {
		numberS1, _ := strconv.ParseInt(string(strNumber[j]), 16, 0)
		numberS2, _ := strconv.ParseInt(string(strNumber[k]), 16, 0)

		// Executa a operação XOR
		newAppendNumber := (numberS1 ^ numberS2) % int64(m)
		strNumber = strNumber[1:] + strconv.FormatInt(newAppendNumber, 16)
		byteNumberFirst := (byte(newAppendNumber) & maskLastBits) << 4

		numberS1, _ = strconv.ParseInt(string(strNumber[j]), 16, 0)
		numberS2, _ = strconv.ParseInt(string(strNumber[k]), 16, 0)

		// Executa a operação XOR
		newAppendNumber = (numberS1 ^ numberS2) % int64(m)
		strNumber = strNumber[1:] + strconv.FormatInt(newAppendNumber, 16)
		byteNumberLast := (byte(newAppendNumber) & maskLastBits)

		// Byte aleatório formado (OR)
		byteNumber := byteNumberFirst | byteNumberLast

		// Cria um string com esses bytes hexadecimais gerados
		numberHexaString = numberHexaString + fmt.Sprintf("%02X", byteNumber)
	}

	bigInt := new(big.Int)
	bigInt.SetString(numberHexaString, 16)

	return bigInt
}
