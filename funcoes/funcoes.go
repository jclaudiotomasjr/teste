package main

import "fmt"

func plus(a, b int) int {
	return a + b
}

func plusPlus(a, b, c int) int {
	return a + b + c
}

func main() {
	resp := plus(8, 9)
	fmt.Println("8 + 9=", resp)

	resp = plusPlus(8, 9, 14)
	fmt.Println("8 + 9 + 14= ", resp)

}
