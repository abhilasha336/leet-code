package main

import (
	"fmt"
	"sort"
)

func main() {
	sli := []int{1, 2, 3, 1, 4, 3}
	var subArr []int
	var finalArr []int
	k := 3
	p := k % len(sli)
	for i := 0; i < p; i++ {
		for j := i; j < k+i; j++ {
			subArr = append(subArr, sli[j])
			fmt.Println("", subArr)

		}
		sort.Ints(subArr)
		finalArr = append(finalArr, subArr[len(subArr)-1])
	}
	fmt.Println("final arayt", finalArr)

}
