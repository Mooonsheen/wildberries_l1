// Реализовать паттерн «адаптер» на любом примере.

package main

import "fmt"

// Интерфейс, который необходим для работы клиентского кода
type Target interface {
	Request() string
}

// Адаптируемый класс
type Adaptee struct {
}

func (a *Adaptee) SpecificRequest() string {
	return "Specific request"
}

// Адаптер
type Adapter struct {
	adaptee *Adaptee
}

func (adapter *Adapter) Request() string {
	// Вызываем нужный метод адаптируемого класса
	return adapter.adaptee.SpecificRequest()
}

func main() {
	adaptee := &Adaptee{}
	adapter := &Adapter{adaptee: adaptee}

	// Клиентский код работает через интерфейс Target
	fmt.Println(adapter.Request())
}
