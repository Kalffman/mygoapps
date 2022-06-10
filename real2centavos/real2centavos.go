package main

import (
	"fmt"
	"os"
	"strconv"
)

func coinsAndRemainder(coin int, cents int) (int, int) {
	return cents / coin, cents % coin
}

func main() {
	for _, sReais := range os.Args[1:] {
		fReais, _ := strconv.ParseFloat(sReais, 64)

		cents := int(fReais * 100)

		qtdCent_50, cents := coinsAndRemainder(50, cents) // 25 centavos

		qtdCent_25, cents := coinsAndRemainder(25, cents) // 25 centavos

		qtdCent_10, cents := coinsAndRemainder(10, cents) // 10 centavos

		qtdCent_5, cents := coinsAndRemainder(5, cents) // 5 centavos

		qtdCent_1, cents := coinsAndRemainder(1, cents) // 1 centavos

		fmt.Printf("%d coins -> 50 centavos\n", qtdCent_50)
		fmt.Printf("%d coins -> 25 centavos\n", qtdCent_25)
		fmt.Printf("%d coins -> 10 centavos\n", qtdCent_10)
		fmt.Printf("%d coins -> 5 centavos\n", qtdCent_5)
		fmt.Printf("%d coins -> 1 centavos\n", qtdCent_1)
	}
}
