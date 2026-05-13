package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("Please provide filename")
		return
	}

	file := args[1]

	data, err := os.ReadFile(file)

	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	lines := strings.Split(string(data), "\n")

	nums := []int{}

	for _, line := range lines {

		line = strings.TrimSpace(line)

		if line == "" {
			continue
		}

		num, err := strconv.Atoi(line)

		if err != nil {
			fmt.Println("Invalid number:", line)
			continue
		}
		nums = append(nums, num)

	}

	if len(nums) == 0 {
		fmt.Println("No valid numbers")
		return
	}

	sum := 0
	max := nums[0]
	min := nums[0]

	for _, num := range nums {

		if num > max {
			max = num
		}

		if num < min {
			min = num
		}

		sum += num
	}

	average := float64(sum) / float64(len(nums))

	fmt.Println("Numbers:", nums)
	fmt.Println("Max:", max)
	fmt.Println("Min:", min)
	fmt.Println("Sum:", sum)
	fmt.Println("Average:", average)

}
