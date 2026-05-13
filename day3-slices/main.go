package main

import "fmt"

func main() {

	var n int

	fmt.Print("Enter number of elements: ")
	fmt.Scan(&n)

	nums := []int{}

	for i := 0; i < n; i++ {

		var x int
		fmt.Scan(&x)

		nums = append(nums, x)
	}

	if len(nums) == 0 {
		fmt.Println("No numbers entered")
		return
	}

	max := nums[0]
	min := nums[0]

	for i := 1; i < len(nums); i++ {

		if nums[i] > max {
			max = nums[i]
		}

		if nums[i] < min {
			min = nums[i]
		}
	}

	fmt.Println("Numbers:", nums)
	fmt.Println("Max:", max)
	fmt.Println("Min:", min)
}
