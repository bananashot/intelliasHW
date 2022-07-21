package main

import (
	"fmt"
	"sync"
)

// Конкурентно порахувати суму кожного слайсу int, та роздрукувати результат.
// Потрібно використовувати WaitGroup.
// Приклад:
// [ [ 4, 6 ], [ 7, 9 ] ]
// Результат друку:
// Порядок друку не важливий.
// “slice 1: 10”
// “slice 2: 16”

func main() {

	n := [][]int{
		{2, 6, 9, 24},
		{7, 3, 94, 3, 0},
		{4, 2, 8, 35},
	}

	// Ваша реалізація
	var wg sync.WaitGroup

	for i, v := range n {
		wg.Add(1)
		go func(n int, slice []int) {
			defer wg.Done()
			fmt.Printf("\nSlice %v: %v", n, sum(slice))
		}(i, v)
	}
	wg.Wait()
}

func sum(slice []int) int {
	result := 0

	for _, v := range slice {
		result += v
	}

	return result
}
