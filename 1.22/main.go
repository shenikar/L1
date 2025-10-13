package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Введите число a: ")
	aStr, _ := reader.ReadString('\n')
	aStr = strings.TrimSpace(aStr)

	fmt.Println("Введите число b: ")
	bStr, _ := reader.ReadString('\n')
	bStr = strings.TrimSpace(bStr)

	// Создаем big.Int и читаем из строк
	a := new(big.Int)
	b := new(big.Int)
	_, okA := a.SetString(aStr, 10)
	_, okB := b.SetString(bStr, 10)
	if !okA || !okB {
		fmt.Println("Ошибка: Введите корректные числа.")
		return
	}

	// Арифметические операции
	sum := new(big.Int).Add(a, b)
	sub := new(big.Int).Sub(a, b)
	mul := new(big.Int).Mul(a, b)

	div := new(big.Int)
	if b.Cmp(big.NewInt(0)) != 0 {
		div.Div(a, b)
	} else {
		div = nil
	}

	//Результаты
	fmt.Println("\nРезультаты операций:")
	fmt.Println("a + b =", sum)
	fmt.Println("a - b =", sub)
	fmt.Println("a * b =", mul)

	if div != nil {
		fmt.Println("a / b =", div)
	} else {
		fmt.Println("Ошибка: Деление на ноль")
	}
}
