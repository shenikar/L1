package main

import "fmt"

// generator отправляет числа из массива в канал.
func generator(nums []int, out chan<- int) {
	defer close(out)
	for _, num := range nums {
		out <- num
	}
}

// doubler читает числа из входного канала, удваивает их и отправляет в выходной канал.
func doubler(in <-chan int, out chan<- int) {
	defer close(out)
	for num := range in {
		out <- num * 2
	}
}

func main() {
	numbers := []int{1, 2, 3, 4, 5}

	naturals := make(chan int)
	doubles := make(chan int)

	go generator(numbers, naturals)
	go doubler(naturals, doubles)

	// Читаем из второго канала и выводим в stdout.
	for result := range doubles {
		fmt.Println(result)
	}
}
