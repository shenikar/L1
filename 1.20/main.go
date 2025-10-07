package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите строку: ")
	line, _ := reader.ReadString('\n')
	line = strings.TrimSpace(line)

	reversed := reverseWord(line)
	fmt.Println("Результат:", reversed)
}

// reverseWord переворачивает порядок слов в строке
func reverseWord(s string) string {
	r := []rune(s)

	//переворачиваем всю строку
	reverseRunes(r, 0, len(r)-1)

	//переворачиваем слова внутри строки
	start := 0
	for i := 0; i <= len(r); i++ {
		if i == len(r) || r[i] == ' ' {
			reverseRunes(r, start, i-1)
			start = i + 1
		}
	}
	return string(r)
}

// reverseRunes переворачивает подмассив рун
func reverseRunes(runes []rune, left, right int) {
	for left < right {
		runes[left], runes[right] = runes[right], runes[left]
		left++
		right--
	}
}
