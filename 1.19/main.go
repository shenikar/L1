package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	input := bufio.NewReader(os.Stdin)
	fmt.Print("Введите строку: ")
	line, err := input.ReadString('\n')
	if err != nil {
		fmt.Fprintln(os.Stderr, "Ошибка чтения:", err)
		return
	}
	line = strings.TrimRight(line, "\r\n")

	reversed := reverseWord(line)
	fmt.Println("Перевернутая строка:", reversed)
}

func reverseWord(word string) string {
	r := []rune(word)
	i, j := 0, len(r)-1
	for i < j {
		r[i], r[j] = r[j], r[i] // меняем местами
		i++
		j--
	}
	return string(r)
}
