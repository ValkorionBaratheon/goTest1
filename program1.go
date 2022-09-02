package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// a. Takes an integer from the command line.
	fmt.Println("Enter an integer: \n")
	input, _ := strconv.Atoi(os.Args[1])

	// b. Store integers 0 to x to Array, Slice, Map
	var arr []int
	sl := make([]int, input)
	m := make(map[int]int)

	for i := 0; i < input; i++ {
		arr[i] = i
		sl[i] = i
		m[i] = i
	}

	fmt.Println(arr)
	os.Exit(1)
	// c. Output time to initialize each data type

	// d. Output time it takes to increment each element in each data type by 1
}
