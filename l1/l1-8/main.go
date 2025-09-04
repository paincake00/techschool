package main

import (
	"fmt"
	"os"
)

func main() {
	// вводимое число
	var d int64
	fmt.Println("Enter your number (int64): ")
	_, err := fmt.Fscan(os.Stdin, &d)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	// индекс бита, который нудно поменять (начиная с 1)
	var i int8
	fmt.Println("Enter your index (1-63): ")
	_, err = fmt.Fscan(os.Stdin, &i)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	// изменяем доступные биты (без бита знака)
	if i < 1 || i > 63 {
		fmt.Println("index only between 1 and 63")
		return
	}
	// определим что на что менять: 1 или 0
	var isOne bool
	fmt.Println("Enter true (change on 1) or false (change on 0): ")
	_, err = fmt.Fscan(os.Stdin, &isOne)
	if err != nil {
		fmt.Println("Error reading input:", err)
	}

	// смещаем бит на нужную позицию (счет с 1)
	ib := int64(1) << (i - 1)

	if isOne {
		// при замене на 1 просто используем OR
		d = d | ib
	} else {
		// при замене на 0 используем NOT для ib, а потом AND
		d = d & (^ib)
	}
	// ответ в десятичной и двоичной системе счисления
	fmt.Printf("%d (%b)\n", d, d)
}
