package main

import "fmt"

func main() {
	words := []string{"cat", "cat", "dog", "cat", "tree"}

	makeSet(words)
}

func makeSet(words []string) {
	set := make(map[string]struct{})

	// кладём все элементы в map
	for _, word := range words {
		set[word] = struct{}{}
	}

	// вывод множества
	fmt.Print("Множество = {")
	first := true
	for k := range set {
		if !first {
			fmt.Print(", ")
		}
		fmt.Print(k)
		first = false
	}
	fmt.Println("}")
}
