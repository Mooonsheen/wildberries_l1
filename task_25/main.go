// Реализовать собственную функцию sleep.

package main

import (
	"fmt"
	"time"
)

func sleep(seconds int) {
	<-time.After(time.Second * time.Duration(seconds))
}

func main() {
	var a int
	fmt.Println("Введите количество секунд сна: ")
	fmt.Scan(a)
	sleep(a)
	fmt.Println("End")
}
