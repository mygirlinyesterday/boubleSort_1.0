package main

import "fmt"

func main() {
	var numbers []int = []int{5, 4, 2, 3, 1, 0}
	fmt.Println("Stack is: ", numbers)

	for i := 1; i < len(numbers); i++ {
		sweep(numbers)
		fmt.Printf("Round %d sweeping: %d\n", i, numbers)
	}

}

func sweep(numbers []int) {
	N := len(numbers)
	_1stIndex := 0
	_2ndIndex := 1

	for _2ndIndex < N {
		_1stNum := numbers[_1stIndex]
		_2ndNum := numbers[_2ndIndex]
		//fmt.Println("Sorted: ", _1stNum, _2ndNum)
		if _1stNum > _2ndNum {
			// Swap numbers
			numbers[_1stIndex] = _2ndNum
			numbers[_2ndIndex] = _1stNum
		}

		_1stIndex++
		_2ndIndex++
	}

}
