package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

// Функция воркер
func worker(id int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		fmt.Printf("Worker %d обработал: %d\n", id, job)
		time.Sleep(time.Millisecond * 100)
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

	jobs := make(chan int, 100) // канал для передачи данных
	var wg sync.WaitGroup

	// создаем воркеров
	for i := 1; i <= workers; i++ {
		wg.Add(1)
		go worker(i, jobs, &wg)
	}

	// главная горутина постоянная запись в канал
	go func() {
		for i := 1; i <= 20; i++ {
			jobs <- i
			fmt.Printf("Главная горутина записала: %d\n", i)
			time.Sleep(50 * time.Millisecond)
		}
		close(jobs) // закрываем канал после всех заданий
	}()

	// Ждем завершения всех воркеров
	wg.Wait()
	fmt.Println("Все задания обработаны")
}
