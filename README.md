# Wave Pattern Array Rearrangement

A Go program that rearranges an array into a "wave" pattern according to specific rules. The program takes an array of integers and a value x (where x≥1) and rearranges elements into blocks with specific peak patterns.

## Problem Description

The program rearranges an array following these rules:

1. **Partitioning**: Array is divided into consecutive blocks of 2x+1 elements
2. **Wave Pattern in each block**: For every block (indices 0 to 2x):
   - Element at index x must be the maximum in the block
   - Elements to the left (indices 0 to x-1) must be in non-decreasing order
   - Elements to the right (indices x+1 to 2x) must be in non-increasing order

## Prerequisites

- Go (1.16 or higher)

## Installation

1. Clone the repository:
```bash
git clone https://github.com/Saad-777/question2.git
```

2. Navigate to the project directory:
```bash
cd question2
```

## Usage

1. Run the program:
```bash
go run question2.go
```

2. When prompted:
   - Enter array elements separated by spaces
   - Enter the value of x (must be ≥ 1)

## Example

Input:
```
Enter the array elements separated by spaces:
3 6 5 10 7
Enter the value of x (must be ≥ 1): 2
```

Output:
```
Input Array: [3 6 5 10 7]
x = 2
Result: [3 6 10 7 5]
Pattern verified ✓
```

## Input Requirements

- Array length must be a multiple of (2x+1)
- x must be greater than or equal to 1
- Array elements must be integers
- Elements must be separated by spaces when input

## Input Validation

The program includes validation for:
- Correct array length based on x value
- Valid integer inputs
- Valid x value (≥ 1)

## Error Handling

The program will:
- Prompt for new input if non-integer values are entered
- Display an error if array length is not valid for the given x
- Allow retrying if invalid input is provided

