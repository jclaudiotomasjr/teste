package main

import "fmt"

func main() {
	numeros := []int{8, 9, 14}
	soma := 0
	for _, num := range numeros {
		soma += num
	}
	fmt.Println("Soma: ", soma)

	for i, num := range numeros {
		if num == 14 {
			fmt.Println("Index:", i)
		}
	}

	kvs := map[string]string{"a": "maça", "b": "laranja"}

	for s, v := range kvs {
		fmt.Printf("%s: -> %s\n", s, v)
	}

	for s := range kvs {
		fmt.Println("Chave:", s)
	}

	for i, c := range "jr" {
		fmt.Println(i, c)
	}

}
