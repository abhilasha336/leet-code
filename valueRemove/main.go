package main

import "fmt"

func removeElement(nums []int, val int) {
	var num1 []int
	for _, value := range nums {
		if value != val {
			num1 = append(num1, value)
		}
	}
	fmt.Println(num1)

}
func removeElement1(nums []int, val int) {
	var num1, num2 []int
	for i, value := range nums {
		if value == val {
			num1 = nums[0:i]
			num2 = nums[i+1:]
			num1 = append(num1, num2...)
		}
	}
	fmt.Println(num1)

}
func main() {
	removeElement([]int{1, 2, 3, 4, 5}, 3)
	removeElement1([]int{1, 2, 3, 4, 5}, 3)

}
