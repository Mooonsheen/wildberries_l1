// Разработать программу, которая будет последовательно отправлять значения в канал,
// а с другой стороны канала — читать. По истечению N секунд программа должна завершаться.

package main

import (
	"fmt"
	"time"
)

// Запись данных в канал
func sendValues(ch chan int) {
	for i := 1; i <= 999999999; i++ {
		ch <- i
		time.Sleep(500 * time.Millisecond)
	}
	close(ch)
}

// Чтение данных из канала
func readValues(ch chan int) {
	for value := range ch {
		fmt.Println(value)
	}
}

func main() {
	var a time.Duration
	fmt.Printf("Введите время работы в секундах: ")
	fmt.Scan(&a)
	ch := make(chan int)
	go sendValues(ch)
	go readValues(ch)

	time.Sleep(a * time.Second)
}
