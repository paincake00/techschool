package main

import (
	"fmt"
	"os"
)

func main() {
	var a, b int
	fmt.Println("Enter a and b: ")
	_, err := fmt.Fscan(os.Stdin, &a, &b)
	if err != nil {
		fmt.Println("Error reading stdin:", err)
		return
	}

	// замена через сложение/вычитание
	a = a + b
	b = a - b
	a = a - b

	fmt.Printf("a = %d, b = %d", a, b)
}
