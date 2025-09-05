package main

import (
	"fmt"
)

func main() {
	arr1 := [...]int{9, 7, 5, 3, 1, 2, 4, 6, 8, 0}
	arr2 := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	arr3 := [...]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	arr4 := [...]int{5, 1, 3, 3, 2, 5, 1, 4, 2, 5}
	arr5 := [...]int{7, 7, 7, 7, 7, 7, 7, 7}
	arr6 := [...]int{42}
	arr7 := [...]int{}
	arr8 := [...]int{10, -1, 3, 7, 2}
	arr9 := [...]int{-3, -1, -7, 4, 2, 0, -2}
	arr10 := [...]int{1000, -500, 200, 0, 999, -1000, 500, -200}

	tests := [10][]int{arr1[:], arr2[:], arr3[:], arr4[:], arr5[:], arr6[:], arr7[:], arr8[:], arr9[:], arr10[:]}

	for _, testArr := range tests {
		fmt.Println(quickSort(testArr))
	}
}

func quickSort(arr []int) []int {
	// нет смысла сортировать если 1 элемент или пусто
	if len(arr) < 2 {
		return arr
	}

	sortOfPart(arr, 0, len(arr)-1)

	return arr
}

func sortOfPart(arr []int, start, end int) {
	// базовый случай рекурсии
	if start >= end {
		return
	}

	// два указателя
	left := start
	right := end

	// для опорного элемента индекс в середине
	mid := (start + end) / 2

	// продолжаем пока указатели не поменяются
	for left < right {
		// ищем элемент больше опорного
		for arr[left] < arr[mid] {
			left++ // добавляем пока не найдем
		}
		// ищем элемент меньше опорного
		for arr[right] > arr[mid] {
			right-- // уменьшаем пока не найдем
		}
		// завершаем если уже пересеклись указатели
		if left > right {
			break
		}
		// свап значений
		arr[left], arr[right] = arr[right], arr[left]
		// двигаем указатели
		left++
		right--
	}

	// применяем для двух половин
	sortOfPart(arr, start, right)
	sortOfPart(arr, left, end)
}
