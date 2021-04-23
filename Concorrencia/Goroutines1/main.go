package main

import(
	"fmt"
	//"time"
)

func imprimir(n int)  {
	for i := 0 ; i < 3; i++ {
		fmt.Printf("%d - ", n)
		time.Sleep(200 * time.Millisecond)
	}
}

func main() {

	//chamada de função Goroutine 
	imprimir(3)
	imprimir(14)

	//Usando Goroutine
	go imprimir(8)
	imprimir(9)
	//time.Sleep(time.Second)
}