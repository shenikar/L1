package main

import "fmt"

func main() {
	num1, num2 := 5, 10
	fmt.Println("До:", num1, num2)

	swap(num1, num2)
	swapXOR(num1, num2)
}

// меняет местами с помощью сложения и вычитания
func swap(num1, num2 int) {
	num1 = num1 + num2
	num2 = num1 - num2
	num1 = num1 - num2

	fmt.Println("После:", num1, num2)
}

// меняет местами с помощью XOR обмена
func swapXOR(num1, num2 int) {
	num1 = num1 ^ num2
	num2 = num1 ^ num2
	num1 = num1 ^ num2

	fmt.Println("После XOR:", num1, num2)
}
