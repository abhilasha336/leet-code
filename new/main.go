package main

import "fmt"

func check(val int, nums []int) []int {
	var indexSlice []int
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == val {
				indexSlice = append(indexSlice, i, j)
				return indexSlice

			}
		}

	}
	return indexSlice

}
func checkOne(val int, nums []int) []int {
	newMap := make(map[int]int)
	var indexSlice []int

	for i, num := range nums {
		complement := val - num
		if j, ok := newMap[complement]; ok {
			indexSlice = append(indexSlice, j, i)
			return indexSlice
		}
		newMap[num] = i
	}
	return nil
}
func main() {

	nums := []int{3, 2, 4}
	sum := 6
	output := checkOne(sum, nums)
	fmt.Println("output", output)

}

// Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

// You may assume that each input would have exactly one solution, and you may not use the same element twice.

// Example 1:

// Input: nums = [2,7,11,15], target = 9
// Output: [0,1]
// Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].

// Example 2:

// Input: nums = [3,2,4], target = 6
// Output: [1,2]

// Example 3:

// Input: nums = [3,3], target = 6
// Output: [0,1]
