package main

import "fmt"

func fibonacci(n int) int {
	if n > 1 {
		return fibonacci(n-1) + fibonacci(n-2)
	} else if n == 1 {
		return 1
	} else if n == 0 {
		return 0
	} else {
		return -1
	}
}

func main() {
	var num int
	fmt.Scanln(&num)
	fmt.Println(fibonacci(num))
}
