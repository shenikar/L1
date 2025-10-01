package main

import "fmt"

func main() {
	detectedType(42)
	detectedType("hello")
	detectedType(true)

	chInt := make(chan int)
	detectedType(chInt)

	chStr := make(chan string)
	detectedType(chStr)

	detectedType(1.1)
}

// определяет тип переменной
func detectedType(i any) {
	switch v := i.(type) {
	case int:
		fmt.Println("Тип int, значение =", v)
	case string:
		fmt.Println("Тип string, значение =", v)
	case bool:
		fmt.Println("Тип bool, значение =", v)
	case chan int:
		fmt.Println("Тип chan int")
	case chan string:
		fmt.Println("Тип chan string")
	default:
		fmt.Println("Неизвестный тип")
	}

}
