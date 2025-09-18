package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("=== Пример 1: Выход по условию ===")
	exitByCondition()
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
