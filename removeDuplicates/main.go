package main

import "fmt"

func removeDuplicates(nums []int) {
	unique := make(map[int]bool)
	var num3 []int
	for _, value := range nums {

		if exists := unique[value]; !exists {
			unique[value] = true
			num3 = append(num3, value)
		}

	}

	fmt.Println(num3)

}
func removeDuplicates1(nums []int) {
	unique := make(map[int]int)
	numss := nums
	// var num3 []int
	count := 1
	for i, value := range nums {

		if value == numss[i] {
			count = count + 1

			unique[value] = count

		}
		count = 0

	}

	fmt.Println("map is", unique)
	var num3 []int
	for key, _ := range unique {
		num3 = append(num3, key)
	}
	fmt.Println(num3)

}
func main() {
	removeDuplicates([]int{1, 2, 2, 3, 3})
	removeDuplicates1([]int{1, 2, 2, 3, 3})

}
