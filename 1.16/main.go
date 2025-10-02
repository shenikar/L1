package main

import "fmt"

func main() {
	nums := []int{2, 5, 3, 4, 10, 9, 22, 15}
	fmt.Println("Исходный массив:", nums)

	sorted := quickSort(nums)
	fmt.Println("Отсортированный массив:", sorted)
}

func quickSort(nums []int) []int {
	var left, right, equal []int

	if len(nums) < 2 {
		return nums
	}

	// опорный эелемент(середина)
	pivot := nums[len(nums)/2]

	for _, v := range nums {
		if v < pivot { // элементы меньше опорного элемента
			left = append(left, v)
		} else if v > pivot { // элементы больше опорного элемента
			right = append(right, v)
		} else { // элементы равные опорному элементу
			equal = append(equal, v)
		}
	}

	// рекурсивная сортировка
	return append(append(quickSort(left), equal...), quickSort(right)...)
}
