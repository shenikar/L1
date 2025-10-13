package main

import "fmt"

func main() {
	slice := []int{10, 20, 30, 40, 50}

	i := 4

	fmt.Println("Исходный слайс:", slice)

	slice = removeIndex(slice, i)
	fmt.Println("После удаления:", slice)

}

// удаляет i элемент по индексу
func removeIndex(slice []int, i int) []int {
	if i < 0 || i >= len(slice) {
		return slice // индекс вне диапазона
	}

	// сдвигаем элементы после i влево
	copy(slice[i:], slice[i+1:])

	// зануляем последний элемент чтобы не было утечки памяти
	slice[len(slice)-1] = 0

	// Возвращаем слайс меньшей длины
	return slice[:len(slice)-1]
}
