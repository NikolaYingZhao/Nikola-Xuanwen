package main

import "fmt"

func swap(a, b *int) {
	*b, *a = *a, *b
}

func anotherswap(a, b int) (int, int) {
	return b, a
}

func main() {
	a, b := 3, 4
	a, b = anotherswap(a, b)
	fmt.Println(a, b)
}
