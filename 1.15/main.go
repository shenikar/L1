package main

var justString string

func main() {
	someFunc()
}

func someFunc() {
	v := createHugeString(1 << 10)
	justString = string([]rune(v)[:100]) // Копируем символы в новый срез байт
}
