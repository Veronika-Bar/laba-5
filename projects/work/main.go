package main

import (
	"fmt"
	"sync"
)

func work() {
	fmt.Println("Работа выполнена")
}

func main() {
	var wg sync.WaitGroup

	// Запускаем 10 горутин
	for i := 0; i < 10; i++ {
		wg.Add(1) // Увеличиваем счетчик горутин
		go func() {
			defer wg.Done() // Уменьшаем счетчик после завершения работы горутины
			work()          // Вызываем функцию work
		}()
	}

	// Ожидаем завершения всех горутин
	wg.Wait()
	fmt.Println("Все работы завершены")
}
