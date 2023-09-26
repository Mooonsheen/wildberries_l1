// Реализовать постоянную запись данных в канал (главный поток). Реализовать
// набор из N воркеров, которые читают произвольные данные из канала и выводят
// в stdout. Необходима возможность выбора количества воркеров при старте.

// Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать способ
// завершения работы всех воркеров.

package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

// Логика горутин
func worker(wg *sync.WaitGroup, ch <-chan int, id int) {
	defer wg.Done()
	for data := range ch {
		fmt.Printf("Воркер %d обработал число %d\n", id, data)
		time.Sleep(500 * time.Microsecond)
	}
}

// Окончание работы горутин и закрытие канала
func monitorWorker(wg *sync.WaitGroup, ch chan int) {
	wg.Wait()
	close(ch)
}

func main() {
	ch := make(chan int)
	var count int
	fmt.Scan(&count) // Количество воркеров

	var wg sync.WaitGroup
	wg.Add(count)

	// Запуск воркеров
	for i := 0; i < count; i++ {
		go worker(&wg, ch, i)
	}

	// Запись данных в канал
	go func() {
		for i := 1; ; i++ {
			ch <- i
		}
	}()

	// Ждем сигнал завершения программы
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// Завершаем работу всех воркеров
	monitorWorker(&wg, ch)
}
