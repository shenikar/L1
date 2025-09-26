package main

import "fmt"

// setBit устанавливает бит в позиции pos числа n в значение value (1 или 0).
func setBit(n int64, pos uint, value int) int64 {
	mask := int64(1 << pos)

	// Проверяем, какое значение нужно установить.
	if value == 1 {
		return n | mask
	}

	// Установка бита в 0 (сброс бита).
	return n &^ mask
}

func main() {
	var num int64 = 5 // 0101 в двоичной системе
	fmt.Printf("Исходное число: %d (%08b)\n", num, num)

	// нумерация битов начинается с 0, но в задании в примере написано
	// то, что для числа 5 (0101₂) установка 1-го бита в 0 даст 4 (0100₂)
	// если нумерацию надо начинать именно с 1, то тогда надо написать res := setBit(num, pos-1, 0)
	pos := uint(0)
	res := setBit(num, pos, 0)
	fmt.Printf("Установка %d-го бита в 0: %d (%08b)\n", pos, res, res)
}
