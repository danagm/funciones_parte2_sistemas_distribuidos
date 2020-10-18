package main

import "fmt"

func mayor(args ...int) int { // args es un slice de los parámetros recibidos
	max := args[0]
	for _, v := range args { // iterar sobre los parámetros como un slice
		if max < v {
			max = v
		}
	}

	return max
}

func main() {
	a := []int{1, 4, 2}
	fmt.Println(mayor(4, 5, 2))
	fmt.Println(mayor(a...)) // cada elemento del slice se envía como parámetros
}
