package main

import "fmt"

func twoSum(numbers []int, target int) (int, int, int) {
	var flag bool = false
	result := make([]int, 2)
	for i := 0; i < len(numbers)-1; i++ {
		for j := i; j < len(numbers)-1; j++ {
			if numbers[i]+numbers[j] == target {
				result[0] = numbers[i]
				result[1] = numbers[j]
				flag = true
			}
		}
	}

	if flag {
		return target, result[0], result[1]
	} else {
		fmt.Println("No valid pair exists for", target)
		panic("No valid pair")
	}
}

func main() {
	data := []int{1, 7, 3, 4, 8, 9}
	var x int = 55

	target, result1, result2 := twoSum(data, x)
	fmt.Println("Pair with a given sum", target, "is (", result1, ",", result2, ")")
}
