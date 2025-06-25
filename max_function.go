package main

import "fmt"

func takeMax(nums ...int) int {
	if len(nums) == 0 {
		return 0
	}
	maxVal := nums[0] 
	for _, currentNumber := range nums[1:] {
		if currentNumber > maxVal {
			maxVal = currentNumber
		} else {
			continue
		}
	}

	return maxVal
}

func main() {
	fmt.Println(takeMax(100, 2, 3, 4, 5, 5))     // 100
	fmt.Println(takeMax(-5, -2, -10))            // -2 ✅
	fmt.Println(takeMax())                       // 0
}
