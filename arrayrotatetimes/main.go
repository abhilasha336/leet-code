package main

import "fmt"

func main() {
	rotate([]int{1, 2, 3, 4, 5, 6, 7}, 3)
}
func rotate(nums []int, k int) {
	n := len(nums)
	k = k % n
	fmt.Println("k", k, n)
	for i := 0; i < k; i++ {
		temp := nums[n-1]
		for j := n - 1; j > 0; j-- {
			nums[j] = nums[j-1]

		}
		nums[0] = temp
		fmt.Println("rotated:", nums)
	}
}
