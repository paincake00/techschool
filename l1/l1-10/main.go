package main

import (
	"fmt"
	"math"
)

func main() {
	// последовательность температур
	temps := [...]float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	// по ключу будем хранить множество
	m := map[int][]float64{}

	for _, t := range temps {
		// получаем ключ-диапазон с помощью отброса дробной части и div
		k := int(math.Trunc(t)) / 10 * 10
		if val, ok := m[k]; ok {
			// если диапозон есть, обновляем множество
			m[k] = append(val, t)
		} else {
			// иначе создаем диапозон с множеством
			m[k] = []float64{t}
		}
	}

	// красивый вывод с форматированием
	for k, v := range m {
		fmt.Printf("%d:{", k)
		for i, n := range v {
			if i > 0 {
				fmt.Print(", ")
			}
			fmt.Printf("%.1f", n)
		}
		fmt.Print("}\n")
	}
}
