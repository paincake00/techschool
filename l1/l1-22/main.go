package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strings"
)

func inputHandler(a, b *big.Int) {
	reader := bufio.NewReader(os.Stdin)

	// первая переменная
	fmt.Print("Input first big number (enter for skip): ")
	line1, _ := reader.ReadString('\n')
	line1 = strings.TrimSpace(line1)
	if line1 != "" {
		_, ok := a.SetString(line1, 10)
		if !ok {
			fmt.Println("Invalid input, using default value")
		}
	}

	// вторая переменная
	fmt.Print("Input second big number (enter for skip): ")
	line2, _ := reader.ReadString('\n')
	line2 = strings.TrimSpace(line2)
	if line2 != "" {
		_, ok := b.SetString(line2, 10)
		if !ok {
			fmt.Println("Invalid input, using default value")
		}
	}
}
func main() {
	// переменные со значениями по умолчанию
	a := new(big.Int).Exp(big.NewInt(2), big.NewInt(20), nil)
	b := new(big.Int).Exp(big.NewInt(2), big.NewInt(21), nil)

	inputHandler(a, b)

	var res big.Int
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
	// сложение
	fmt.Println("sum: ", res.Add(a, b))
	// умножение
	fmt.Println("mul: ", res.Mul(a, b))
	// вычитание
	fmt.Println("sub: ", res.Sub(a, b))
	// деление с плавающей точкой
	fmt.Println("div: ", new(big.Float).Quo(new(big.Float).SetInt(a), new(big.Float).SetInt(b)))
}
