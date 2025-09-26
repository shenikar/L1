package main

import "fmt"

// Родительская структура
type Human struct {
	Name string
	Surname string
	Age int
}

// Метод структуры Human
func (h Human) Walk() {
	fmt.Println(h.Name, "is walking")
}

// Дочерняя структура, наследующая Human
type Action struct {
	Human // Встраивание структуры Human
}

func main() {
	// Создание экземпляра Action
	action := Action{
		Human: Human{Name: "Alex", Surname: "Smith", Age: 29}}
	action.Walk()
	}
