package main

import "fmt"

func main() {
	fmt.Println(twoSum167([]int{2, 7, 11, 15}, 9))
}

func twoSum167(numbers []int, target int) []int {
	i, j := 0, len(numbers)-1
	for i < j {
		sum := numbers[i] + numbers[j]
		if sum == target {
			return []int{i+1, j+1}
		}
		if sum < target {
			i += 1
		} else {
			j -= 1
		}
	}
	return []int{}
}
