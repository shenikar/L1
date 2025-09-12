package main

import (
	"fmt"
	"time"
)

func main() {
	countTime := 5 // секунды до завершения программы

	data := make(chan int)

	// Писатель (генератор чисел)
	go func() {
		i := 1
		for {
			data <- i
			i++
			time.Sleep(500 * time.Millisecond)
		}

	}()

	// Читатель (основная горутина)
	timeout := time.After(time.Duration(countTime) * time.Second)

	for {
		select {
		case val := <-data:
			fmt.Println("Прочитано из канала:", val)
		case <-timeout:
			fmt.Println("Время вышло, завершаем программу.")
			return
		}
	}
}
