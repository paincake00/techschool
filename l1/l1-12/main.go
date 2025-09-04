package main

import (
	"fmt"
)

func main() {
	words := []string{"cat", "cat", "dog", "cat", "tree"}

	fmt.Println(properSet(words))
}

func properSet(words []string) []string {
	set := make(map[string]struct{})
	var res []string

	for _, w := range words {
		if _, ok := set[w]; !ok { // элемент новый
			set[w] = struct{}{}
			// добавляем в итоговое множество
			res = append(res, w)
		}
	}

	return res
}
