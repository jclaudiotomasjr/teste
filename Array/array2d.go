package main

import "fmt"

func main() {

	var array1 [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			array1[i][j] = i + j
		}
	}
	fmt.Println("Array: ", array1)
}