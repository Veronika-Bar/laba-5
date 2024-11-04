package main

import "fmt"

// calculator принимает два канала чисел и один канал для остановки,
// а затем возвращает канал с обработанными значениями.
func calculator(firstChan <-chan int, secondChan <-chan int, stopChan <-chan struct{}) <-chan int {
	outChan := make(chan int)

	go func() {
		defer close(outChan)

		for {
			select {
			case val := <-firstChan:
				outChan <- val * val // отправляем квадрат значения
			case val := <-secondChan:
				outChan <- val * 3 // отправляем значение, умноженное на 3
			case <-stopChan:
				return // выход из функции при получении сигнала остановки
			}
		}
	}()

	return outChan
}

func main() {
	firstChan := make(chan int)
	secondChan := make(chan int)
	stopChan := make(chan struct{})

	outChan := calculator(firstChan, secondChan, stopChan)

	// Пример использования
	go func() {
		firstChan <- 6
		secondChan <- 1
		stopChan <- struct{}{}
	}()

	for result := range outChan {
		fmt.Println(result)
	}
}
