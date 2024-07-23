package main

import "fmt"

//after twicw element occured , 3rd repeated elements should be removed
func removeElement(nums []int) {
	// var num1 []int
	count := 0
	insertPosition := 0
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			count = count + 1
		} else {
			count = 0
		}
		if count <= 2 {
			nums[insertPosition] = nums[i+1]
			insertPosition++
		}
	}
	fmt.Println(nums)

}

func main() {
	removeElement([]int{1, 1, 1, 2, 2, 3})

}
