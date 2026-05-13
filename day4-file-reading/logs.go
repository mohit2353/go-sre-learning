package main
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

    counts := make(map[string]int)

    for _, line := range lines {

        line = strings.TrimSpace(line)

        if line == "" {
            continue
        }

        parts := strings.Split(line, " ")

        if len(parts) != 2 {
            fmt.Println("Invalid line:", line)
            continue
        }

        logType := parts[0]
        value := parts[1]

        num, err := strconv.Atoi(value)
        if err != nil {
            fmt.Println("Invalid number:", value)
            continue
        }

        nums = append(nums, num)
        counts[logType]++
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

    fmt.Println("Max:", max)
    fmt.Println("Min:", min)

    fmt.Println()

    fmt.Println("INFO count:", counts["INFO"])
    fmt.Println("ERROR count:", counts["ERROR"])
    fmt.Println("WARN count:", counts["WARN"])
}