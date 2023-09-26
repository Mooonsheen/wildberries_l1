// Реализовать все возможные способы остановки выполнения горутины.

package main

import (
	"fmt"
	"time"
)

func myGoroutine(stop chan bool) {
	for {
		select {
		case <-stop:
			fmt.Println("Горутина остановлена")
			return
		default:
			fmt.Println("Выполняются операции")
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	stop := make(chan bool)
	go myGoroutine(stop)

	time.Sleep(5 * time.Second)
	stop <- true

	time.Sleep(2 * time.Second)
	fmt.Println("Завершение программы")
}

// Второй способ
//
// import (
//  "context"
//  "fmt"
//  "time"
// )

// func myGoroutine(ctx context.Context) {
//  for {
//   select {
//   case <-ctx.Done():
//    fmt.Println("Горутина остановлена")
//    return
//   default:
//    fmt.Println("Выполняются операции")
//    time.Sleep(1 * time.Second)
//   }
//  }
// }

// func main() {
//  ctx, cancel := context.WithCancel(context.Background())
//  go myGoroutine(ctx)

//  time.Sleep(5 * time.Second)
//  cancel()

//  time.Sleep(2 * time.Second)
//  fmt.Println("Завершение программы")
// }

// Третий способ
//
// import (
//     "fmt"
//     "sync"
//     "time"
//    )

//    var stopFlag bool
//    var wg sync.WaitGroup

//    func myGoroutine() {
//     defer wg.Done()
//     for {
//      if stopFlag {
//       fmt.Println("Горутина остановлена")
//       return
//      }
//      fmt.Println("Выполняются операции")
//      time.Sleep(1 * time.Second)
//     }
//    }

//    func main() {
//     wg.Add(1)
//     go myGoroutine()

//     time.Sleep(5 * time.Second)
//     stopFlag = true

//     wg.Wait()
//     fmt.Println("Завершение программы")
//    }
