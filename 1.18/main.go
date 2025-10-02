package main

import (
	"fmt"
	"sync"
)

// Counter — потокобезопасный счётчик.
// Использует sync.Mutex для защиты от одновременного доступа из нескольких горутин.
type Counter struct {
	mu    sync.Mutex
	value int
}

func main() {
	var wg sync.WaitGroup

	counter := &Counter{}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			wg.Done()
			counter.Increment()
		}()
	}

	wg.Wait()
	fmt.Println("Итоговое значение счетчика:", counter.Value())
}

// Инкрементирует значение
func (c *Counter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

// Возвращает текущее значение счётчика.
func (c *Counter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}
