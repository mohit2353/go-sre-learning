import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	data, err := os.ReadFile("data.txt")
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
		fmt.Println("No valid data")
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
	fmt.Println("Max is:", max)
	fmt.Println("Min is:", min)
}