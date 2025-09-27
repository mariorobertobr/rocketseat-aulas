package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Jogo de adivinhação")
	fmt.Println("Digite um número entre 1 e 100")

	x := rand.Int63n(101)

	scanner := bufio.NewScanner(os.Stdin)
	chutes := [10]int64{}

	for i := range chutes {

		fmt.Println("Digite o chute")
		scanner.Scan()
		chute := scanner.Text()
		chute = strings.TrimSpace(chute)

		chuteInt, err := strconv.ParseInt(chute, 10, 64)
		if err != nil {
			fmt.Println("Digite um número válido")
			return
		}
		switch {
		case chuteInt > x:
			fmt.Println("O número é menor", chuteInt)
		case chuteInt < x:
			fmt.Println("O número é maior", chuteInt)
		case chuteInt == x:
			fmt.Println("Você acertou!", chuteInt)
			return
		}

		chutes[i] = chuteInt
	}

	fmt.Printf("você não acertou o numero que era: %d \n "+
		", essas foram as tentativas: %v", x, chutes)

}
