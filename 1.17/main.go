package main

import "fmt"

func main() {
	target := 6
	arr := []int{1, 2, 3, 4, 5, 6}

	index := binarySearch(arr, target)
	fmt.Println("Индекс искомого числа: ", index)
}

func binarySearch(arr []int, target int) int {
	first := 0
	last := len(arr) - 1

	for first <= last {
		mid := first + (last-first)/2
		if arr[mid] == target { // нашли сразу элемент
			return mid
		}
		if arr[mid] < target { //если меньше значит ищем справа
			first = mid + 1
		} else { // если больше значит ещем слева
			last = mid - 1
		}
	}
	return -1
}
