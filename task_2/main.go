// Написать программу, которая конкурентно рассчитает значение квадратов
// чисел взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.

package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
)

// Логика горутин
func worker(wg *sync.WaitGroup, ch chan int, n int) {
	defer wg.Done()
	sqr := n * n
	ch <- (sqr)
}

// Окончание работы горутин и закрытие канала
func monitorWorker(wg *sync.WaitGroup, ch chan int) {
	wg.Wait()
	close(ch)
}

func main() {
	nums := []int{2, 4, 6, 8, 10}
	wg := sync.WaitGroup{}
	ch := make(chan int)

	out := bufio.NewWriter(os.Stdout)

	for _, num := range nums[:] {
		wg.Add(1)
		go worker(&wg, ch, num)
	}

	for i := 0; i < 5; i++ {
		var a int = <-ch
		fmt.Fprintf(out, "%d\n", a)
	}

	out.Flush()

	monitorWorker(&wg, ch)
}
