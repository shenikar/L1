package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Start:", time.Now())
	Sleep(5 * time.Second)
	fmt.Println("End:", time.Now())
}

// Sleep блокирует выполнение на заданное время, создавая внутренний таймер в отдельной горутине
func Sleep(d time.Duration) {
	done := make(chan struct{}) // канал для уведомления о завершении таймера

	// Запускаем горутину, которая засыпает на время d
	go func() {
		time.AfterFunc(d, func() { close(done) }) // по истечении d закрываем канал
	}()

	// Блокируемся, пока канал не закроется
	<-done
}

// Sleep блокирует выполнение активным ожиданием (busy-wait)
func Sleep2(d time.Duration) {
	start := time.Now()

	for time.Since(start) < d {
		// ничего не делаем
	}
}
