package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
	"time"
)

// Функция воркер
func worker(ctx context.Context, id int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done(): // сигнал завершения
			fmt.Printf("Worker %d завершает работу\n", id)
			return
		case job, ok := <-jobs:
			if !ok {
				fmt.Printf("Worker %d: канал закрыт\n", id)
				return
			}
			fmt.Printf("Worker %d обработал: %d\n", id, job)
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Использование: go run main.go <кол-во воркеров>")
		return
	}

	// читаем количество воркеров из аргумента
	workers, err := strconv.Atoi(os.Args[1])
	if err != nil || workers < 1 {
		fmt.Println("Неверное количество воркеров")
		return
	}

	// Контекст с отменой по Ctrl+C (SIGINT/SIGTERM)
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	jobs := make(chan int, 100) // канал для передачи данных
	var wg sync.WaitGroup

	// создаем воркеров
	for i := 1; i <= workers; i++ {
		wg.Add(1)
		go worker(ctx, i, jobs, &wg)
	}

	// главная горутина постоянная запись в канал
	go func() {
		i := 1
		for {
			select {
			case <-ctx.Done():
				close(jobs)
				return
			default:
				jobs <- i
				fmt.Printf("Главная горутина записала: %d\n", i)
				i++
				time.Sleep(100 * time.Millisecond)
			}
		}
	}()

	// Ждем завершения всех воркеров
	wg.Wait()
	fmt.Println("Все воркеры завершили работу. Программа закрыта.")
}

// Я выбрал context.Context, потому что:
// 1. это стандартный и идиоматичный инструмент в Go;
// 2. легко комбинируется с таймаутами (context.WithTimeout) и дедлайнами (context.WithDeadline);
// 3. удобно передавать дальше по коду (воркерам, сервисам, в HTTP-запросы и т.д.);
// 4. обеспечивает единый механизм отмены для всего дерева горутин.
