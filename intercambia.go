package main

import "fmt"

func exchange(a, b *int) {
	*a, *b = *b, *a
}

func main() {
	var a int
	var b int
	fmt.Scanln(&a)
	fmt.Scanln(&b)
	exchange(&a, &b)
	fmt.Println(a, b)
}
