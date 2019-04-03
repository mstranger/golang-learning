// p. 177, task 5.15

package main

// TODO: task 5.16 - 5.17

import "fmt"

func min(vals ...int) (error, int) {
	if len(vals) < 1 {
		return fmt.Errorf("Error: no one number given."), 0
	}

	result := vals[0]
	for _, val := range vals[1:] {
		if val < result {
			result = val
		}
	}

	return nil, result
}

func max(vals ...int) (error, int) {
	if len(vals) < 1 {
		return fmt.Errorf("Error: no one number given."), 0
	}

	result := vals[0]
	for _, val := range vals[1:] {
		if val > result {
			result = val
		}
	}

	return nil, result
}

func main() {
	arr := []int{3, 4, 5, 1, 2}
	fmt.Println(min(arr...))
	fmt.Println(min())
	fmt.Println(max(arr...))
	fmt.Println(max())
}
