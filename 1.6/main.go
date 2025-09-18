package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Println("=== Пример 1: Выход по условию ===")
	exitByCondition()
	time.Sleep(2 * time.Second)

	fmt.Println("\n=== Пример 2: Выход по таймеру ===")
	exitByTimer()
	time.Sleep(2 * time.Second)

	fmt.Println("\n=== Пример 3: Выход по контексту ===")
	exitByContext()
	time.Sleep(2 * time.Second)

	fmt.Println("\n=== Пример 4: Выход по каналу ===")
	exitByChannel()
	time.Sleep(2 * time.Second)

}

// 1. Выход по условию
func exitByCondition() {
	go func() {
		i := 0
		for {
			if i >= 5 {
				fmt.Println("Горутина завершается по условию")
				return
			}
			fmt.Println("Горутина работает:", i)
			i++
			time.Sleep(300 * time.Millisecond)
		}
	}()
	time.Sleep(1 * time.Second)
}

// 2. Выход по таймеру
func exitByTimer() {
	stop := time.After(2 * time.Second)
	go func() {
		for {
			select {
			case <-stop: // получаем сигнал остановки
				fmt.Println("Горутина завершается по таймеру")
				return
			default:
				fmt.Println("Горутина работает")
				time.Sleep(300 * time.Millisecond)
			}
		}
	}()
	time.Sleep(2 * time.Second)
}

// 3. Выход по контексту
func exitByContext() {
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		for {
			select {
			case <-ctx.Done(): // получаем сигнал остановки
				fmt.Println("Горутина завершается по контексту")
				return
			default:
				fmt.Println("Горутина работает")
				time.Sleep(300 * time.Millisecond)
			}
		}
	}()

	time.Sleep(1 * time.Second)
	cancel() // вызываем отмену контекста
	time.Sleep(1 * time.Second)
}

// 4. Выход по каналу
func exitByChannel() {
	stop := make(chan struct{})
	go func() {
		for {
			select {
			case <-stop: // получаем сигнал остановки
				fmt.Println("Горутина завершается по каналу")
				return
			default:
				fmt.Println("Горутина работает")
				time.Sleep(300 * time.Millisecond)
			}
		}
	}()

	time.Sleep(1 * time.Second)
	close(stop)
	time.Sleep(1 * time.Second)
}

// 5. Выход по runtime.Goexit
// Этот метод завершает только текущую горутину, не влияя на другие.
func exitByGoexit() {
	go func() {
		fmt.Println("Горутина завершает себя через Goexit")
		runtime.Goexit()
		fmt.Println("Этот код не будет выполнен")

	}()
	time.Sleep(1 * time.Second)
}
