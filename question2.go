package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// rearrangeWave rearranges the array into wave pattern blocks
func rearrangeWave(arr []int, x int) []int {
	blockSize := 2*x + 1
	if len(arr)%blockSize != 0 {
		panic("Array length must be a multiple of 2x+1")
	}

	result := make([]int, len(arr))
	copy(result, arr)

	// Process each block
	for i := 0; i < len(arr); i += blockSize {
		rearrangeBlock(result[i:i+blockSize], x)
	}

	return result
}

// rearrangeBlock rearranges a single block according to the wave pattern
func rearrangeBlock(block []int, x int) {
	// Find the maximum element in the block
	maxVal := block[0]
	for _, val := range block {
		if val > maxVal {
			maxVal = val
		}
	}

	// Create a temporary slice to store the sorted block
	temp := make([]int, len(block))
	copy(temp, block)
	sort.Ints(temp)

	// Place the maximum element at index x
	block[x] = maxVal

	// Fill left side in non-decreasing order
	leftElements := temp[:x]
	copy(block[:x], leftElements)

	// Fill right side in non-increasing order
	rightElements := temp[x : len(temp)-1] // Exclude the max element
	for i := 0; i < x; i++ {
		block[x+1+i] = rightElements[len(rightElements)-1-i]
	}
}

// Function to get array input from user
func getArrayInput() []int {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the array elements separated by spaces:")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	// Split the input string into numbers
	numStrings := strings.Split(input, " ")
	numbers := make([]int, 0)

	// Convert each string to integer
	for _, numStr := range numStrings {
		num, err := strconv.Atoi(strings.TrimSpace(numStr))
		if err != nil {
			fmt.Printf("Invalid input: %s is not a number\n", numStr)
			return getArrayInput() // Recursively ask for input again
		}
		numbers = append(numbers, num)
	}

	return numbers
}

// Function to get x value from user
func getXValue() int {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter the value of x (must be ≥ 1): ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		x, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input: Please enter a number")
			continue
		}

		if x < 1 {
			fmt.Println("x must be greater than or equal to 1")
			continue
		}

		return x
	}
}

func main() {
	// Get input from user
	arr := getArrayInput()
	x := getXValue()

	// Check if array length is valid
	blockSize := 2*x + 1
	if len(arr)%blockSize != 0 {
		fmt.Printf("Error: Array length (%d) must be a multiple of %d (2x+1)\n", len(arr), blockSize)
		return
	}

	fmt.Printf("\nInput Array: %v\n", arr)
	fmt.Printf("x = %d\n", x)

	// Rearrange the array
	result := rearrangeWave(arr, x)
	fmt.Printf("Result: %v\n", result)

	// Verify the result
	if verifyWavePattern(result, x) {
		fmt.Println("Pattern verified ")
	} else {
		fmt.Println("Pattern verification failed ")
	}
}

// verifyWavePattern verifies if the array follows the wave pattern rules
func verifyWavePattern(arr []int, x int) bool {
	blockSize := 2*x + 1

	for i := 0; i < len(arr); i += blockSize {
		block := arr[i : i+blockSize]

		// Check if element at index x is maximum in the block
		for j := 0; j < blockSize; j++ {
			if block[j] > block[x] {
				return false
			}
		}

		// Check left side is non-decreasing
		for j := 1; j < x; j++ {
			if block[j] < block[j-1] {
				return false
			}
		}

		// Check right side is non-increasing
		for j := x + 1; j < blockSize-1; j++ {
			if block[j] < block[j+1] {
				return false
			}
		}
	}

	return true
}
