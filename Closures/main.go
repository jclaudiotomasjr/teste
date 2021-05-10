package main

import "fmt"

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {

	nextInts := intSeq()
	fmt.Println("Var 1: ", nextInts())
	fmt.Println("Var 1: ", nextInts())
	fmt.Println("Var 1: ", nextInts())

	newInts := intSeq()
	fmt.Println("Var 2: ",newInts())

	

}
