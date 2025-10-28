package main

import (
	"fmt"
	"strings"
)

func main() {
	test1 := "abcd"
	test2 := "abCdefAaf"
	test3 := "aAbcd"

	fmt.Println(CountLetter(test1))
	fmt.Println(CountLetter(test2))
	fmt.Println(CountLetter(test3))

	fmt.Println(CountLetterWithMap(test1))
	fmt.Println(CountLetterWithMap(test2))
	fmt.Println(CountLetterWithMap(test3))
}

// CountLetter проверяет, что все символы в строке являются уникальными в нижнем регистре.
// Он использует вложенные циклы для сравнения каждого символа с каждым другим символом.
func CountLetter(str string) bool {
	str = strings.ToLower(str)
	chars := []rune(str)
	lenght := len(chars)
	for i := range lenght {
		for j := i + 1; j < lenght; j++ {
			if chars[j] == chars[i] {
				return false
			}
		}
	}
	return true

}

// CountLetterWithMap проверяет, что все символы в строке уникальны, используя map и не учитывая регистр.
// Он перебирает строку и сохраняет каждый символ в карте. Если символ уже существует на карте, это означает, что он не уникален.
func CountLetterWithMap(str string) bool {
	str = strings.ToLower(str)
	slcstr := strings.Split(str, "")
	res := make(map[string]bool)
	for _, item := range slcstr {
		_, exist := res[item]
		if exist {
			return false
		}
		res[item] = true
	}
	return true
}
