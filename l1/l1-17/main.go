package main

import "fmt"

func main() {
	sorted := []int{0, 0, 2, 4, 5, 7, 10, 11}
	fmt.Println(binarySearch(sorted, 34))
	fmt.Println(binarySearch(sorted, 4))
	fmt.Println(binarySearch(sorted, 11))
}

func binarySearch(s []int, target int) int {
	// можем сразу дать ответ при условиях
	if len(s) == 0 {
		return -1
	} else if len(s) == 1 {
		return 0
	}

	return search(s, target, 0, len(s)-1)
}

func search(s []int, target, left, right int) int {
	// базовый случай рекурсии (если не нашли)
	if left > right {
		return -1
	}

	mid := (left + right) / 2

	if s[mid] > target {
		// берем половину, в которой будет искомый элемент
		return search(s, target, left, mid-1)
	} else if s[mid] < target {
		return search(s, target, mid+1, right)
	}
	// при s[mid]==target
	return mid
}
