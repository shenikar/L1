package main

import (
	"fmt"
	"sync"
)

func main() {
	// массив чисел
	numbers := []int{2, 4, 6, 8, 10}

	//используем WaitGroup чтобы дождаться всех горутин
	var wg sync.WaitGroup

	for _, n := range numbers {
		wg.Add(1) // увеличиваем счетчик горутин
		// запускаем горутину для вычисления квадрата числа
		go func(num int) {
			defer wg.Done() // уменьшаем счетчик при завершении
			square := num * num
			fmt.Printf("Квадрат числа %d равно = %d\n", num, square)
		}(n) // передаем n как аргумент, чтобы не было ошибки захвата переменной 
	}

	// ждем завершения горутин
	wg.Wait()
}
