package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var s string
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter words: ")
	scanner.Scan()
	s = scanner.Text()

	fmt.Println(reverseWords(s))
}

func reverseWords(s string) string {
	// преобразуем строку в срез рун
	runes := []rune(s)

	// разворачиваем срез рун и передаем индексы первой и последней руны
	reverseSlice(runes, 0, len(runes)-1)

	// теперь будем искать слова и уже их обратно разворачивать
	start := 0 // индекс начала нового слова
	for i := 0; i <= len(runes); i++ {
		// перебираем пока не дойдем до конца слова (или среза)
		if i == len(runes) || runes[i] == ' ' {
			// берем начало и конец слова
			reverseSlice(runes, start, i-1)
			// устанавливаем начало следующего слова
			start = i + 1
		}
	}

	return string(runes)
}

// функция для разворота рун в срезе (переиспользуется)
func reverseSlice(slice []rune, start, end int) {
	// перебираем через две переменные, начиная с разных концов среза
	for i, j := start, end; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
}
