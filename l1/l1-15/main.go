package main

import (
	"strings"
)

var justString string

func someFunc() {
	v := createHugeString(1 << 10)
	/*
		justString = v[:100]
		Создается новая строка (v[:100]),
		но она ссылается на те же байты, что и v.
		После завершения функции, пропадает v, но не байты,
		т.к. на них ссылается justString (которая хранит не 10 байтов как ожидалось).
		Это утечка памяти.
		Поэтому пересоздаем строку, копируя только нужные байты.
	*/
	justString = string([]byte(v[:100]))
}

func main() {
	someFunc()
}

func createHugeString(size int) string {
	return strings.Repeat("a", size)
}
