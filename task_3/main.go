// Дана последовательность чисел: 2,4,6,8,10. Найти сумму их квадратов(22+32+42….)
// с использованием конкурентных вычислений.

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
	fmt.Println(sqr)
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
	var sum int

	out := bufio.NewWriter(os.Stdout)

	for _, num := range nums[:] {
		wg.Add(1)
		go worker(&wg, ch, num)
	}

	for i := 0; i < 5; i++ {
		var a int = <-ch
		sum += a
	}

	fmt.Fprintf(out, "%d\n", sum)
	out.Flush()

	monitorWorker(&wg, ch)
}
