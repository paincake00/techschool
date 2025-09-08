package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5}
	i := 2
	s = removeElem(s, i)
	fmt.Println(s)
}

func removeElem(s []int, i int) []int {
	copy(s[i:], s[i+1:])
	return s[:len(s)-i]
}
