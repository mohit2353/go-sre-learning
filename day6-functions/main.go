package main

import "fmt"

func findMax(nums []int) int {

	max := nums[0]
	for _, num := range nums {
		if num > max {
			max = num

		}
	}
	return max
}

func findMin(nums []int) int {

	min := nums[0]
	for _, num := range nums {
		if num < min {
			min = num
		}

	}
	return min
}

func findSum(nums []int) int {

	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

func findAverage(nums []int) float64 {

	sum := findSum(nums)
	average := float64(sum) / float64(len(nums))
	return average
}

func main() {

	nums := []int{3, 5, 2, 8, 1}

	max := findMax(nums)
	min := findMin(nums)
	sum := findSum(nums)
	average := findAverage(nums)

	fmt.Println("Number:", nums)
	fmt.Println("Max:", max)
	fmt.Println("Min:", min)
	fmt.Println("Sum:", sum)
	fmt.Println("Average:", average)

}
