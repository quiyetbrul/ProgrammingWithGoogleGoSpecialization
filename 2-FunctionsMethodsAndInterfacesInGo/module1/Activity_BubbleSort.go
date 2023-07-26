package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func bubbleSort(numbers []int) {
	n := len(numbers)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if numbers[j] > numbers[j+1] {
				swap(numbers, j)
			}
		}
	}
}

func swap(numbers []int, i int) {
	numbers[i], numbers[i+1] = numbers[i+1], numbers[i]
}

func readInputLine() string {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func getUserInput() []int {
	fmt.Println("Enter up to 10 integers (separated by spaces): ")
	input := readInputLine()

	numStrings := strings.Fields(input)
	numbers := make([]int, 0, len(numStrings))
	for _, numStr := range numStrings {
		num, err := strconv.Atoi(numStr)
		if err != nil {
			fmt.Println("Invalid input. Please enter valid integers.")
			return nil
		}
		numbers = append(numbers, num)
	}

	if len(numbers) > 10 {
		fmt.Println("Input contains more than 10 integers. Please try again.")
		return getUserInput()
	}

	return numbers
}

func main() {
	numbers := getUserInput()
	if numbers == nil {
		return
	}

	fmt.Println("Unsorted array:", numbers)
	bubbleSort(numbers)
	fmt.Println("Sorted array:", numbers)
}
