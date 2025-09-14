package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var s string
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter a string: ")
	scanner.Scan()
	s = scanner.Text()

	fmt.Println(checkUniqueness(s))
}

func checkUniqueness(s string) bool {
	// переводим в руны в нижнем регистре
	runes := []rune(strings.ToLower(s))

	// структура для проверки: map
	m := make(map[rune]struct{})

	for _, v := range runes {
		if _, ok := m[v]; ok {
			// если уже есть значение по этому ключу,
			// выводим false, неуникальное
			return false
		}
		// добавляем значение по новому ключу
		m[v] = struct{}{}
	}
	return true
}
