package main

import "fmt"

func main() {
	s1 := []int{2, 1, 3}
	s2 := []int{2, 3, 4}

	fmt.Println(intersection(s1, s2))
}

func intersection(a, b []int) []int {
	// вспомогательная мапа для сохранения элементов из первого множества
	m := map[int]int{}
	// итоговое множетсво
	var out []int

	for _, v := range a {
		// добавляем что-то, по этим ключам будем искать, перебирая второе множество
		m[v] = 0
	}

	for _, v := range b {
		// если элемент из второго множества совпадает с ключом мапы - значит пересечение с первым множетсвом
		if _, ok := m[v]; ok {
			// добавляем в итогое множество
			out = append(out, v)
		}
	}

	return out
}
