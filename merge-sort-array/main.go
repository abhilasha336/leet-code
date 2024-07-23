package main

import (
	"fmt"
	"sort"
)

func merge(nums1 []int, m int, nums2 []int, n int) {

	nums1 = append(nums1, nums2...)
	sort.Ints(nums1)

	fmt.Println(nums1)

}
func merge1(nums1 []int, m int, nums2 []int, n int) {
	var nums3 []int
	nums3 = append(nums3, nums1...)
	nums3 = append(nums3, nums2...)

	sort.Ints(nums3[:m+n])
	fmt.Println(nums3)

}
func main() {
	merge([]int{1, 2, 3}, 3, []int{2, 4, 5}, 3)
	merge1([]int{1, 2, 3}, 3, []int{2, 4, 5}, 3)

}
