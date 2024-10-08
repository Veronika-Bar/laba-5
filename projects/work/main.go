package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func work() {
	time.Sleep(time.Millisecond * 50)
	fmt.Println("done")
}

func main() {
	// Запускаем 10 горутин
	for i := 0; i < 10; i++ {
		wg.Add(1) // Увеличиваем счетчик горутин
		go func() {
			defer wg.Done() // Уменьшаем счетчик после завершения работы горутины
			work()
		}()
	}

	wg.Wait()
}
