package filter

import (
	"fmt"
)

// twoSum returns the indices of the two numbers that add up to target.
func twoSum(nums []int, target int) []int {
	seen := make(map[int]int) // value -> index

	for i, num := range nums {
		diff := target - num

		// Check if the difference already exists in the map
		if idx, ok := seen[diff]; ok {
			return []int{idx, i}
		}

		// Otherwise store the current number's index
		seen[num] = i
	}

	return nil // no solution found
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9

	result := twoSum(nums, target)
	if result != nil {
		fmt.Printf("Indices: %v (values: %d + %d)\n", result, nums[result[0]], nums[result[1]])
	} else {
		fmt.Println("No two numbers add up to target.")
	}
}
