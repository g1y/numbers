package main

import (
	"fmt"
	"math/rand"
)

func main() {
	numbers := scrambledNumbers()
	fmt.Println(numbers)
}

func scrambledNumbers() [10]int {
	numbers := [10]int{}
	for i := 0; i < 10; i++ {
		numbers[i] = i + 1
	}

	swapFunc := func(i, j int) {
		temp := numbers[j]
		numbers[j] = numbers[i]
		numbers[i] = temp
	}
	rand.Shuffle(10, swapFunc)

	return numbers
}
