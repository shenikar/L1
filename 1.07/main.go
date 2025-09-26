package main

import (
	"fmt"
	"sync"
)

// Создаем потокабезопасную структуру для map
type SafeMap struct {
	mu   sync.Mutex
	data map[string]int
}

// NewSafeMap создает новый экземпляр SafeMap
func NewSafeMap() *SafeMap {
	return &SafeMap{
		data: make(map[string]int),
	}
}

// Set устанавливает значение в map
func (sm *SafeMap) Set(key string, val int) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	sm.data[key] = val
}

// Get получает значение из map
func (sm *SafeMap) Get(key string) (int, bool) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	val, ok := sm.data[key]
	return val, ok
}

func main() {
	sm := NewSafeMap()
	var wg sync.WaitGroup

	//Запускаем горутины
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			key := fmt.Sprintf("key %d", n)
			sm.Set(key, n)
		}(i)
	}
	wg.Wait()

	// Выводим количество элементов для проверки.
	fmt.Printf("Map contains %d items\n", len(sm.data))

	for i := 0; i < 5; i++ {
		key := fmt.Sprintf("key %d", i)
		val, ok := sm.Get(key)
		if ok {
			fmt.Printf("%s: %d\n", key, val)
		}
	}
}
