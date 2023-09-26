// Реализовать конкурентную запись данных в map.

package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	m := make(map[int]int)
	mutex := sync.Mutex{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()

			// Ограничиваем доступ остальных горутин на запись по этому ключу
			mutex.Lock()
			m[n] = n * 10
			mutex.Unlock()
		}(i)
	}

	wg.Wait()

	// Выводим значения из карты
	for k, v := range m {
		fmt.Printf("Ключ: %d, Значение: %d\n", k, v)
	}
}
