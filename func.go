package main

import (
	"fmt"
	"reflect"
	"runtime"
)

func calculate(a, b int, op string) (int, error) {
	var q int
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		q, _ = div(a, b)
		return q, nil
	default:
		return 0, fmt.Errorf(
			"unsupported operation:%s", op)
	}
}

func div(a, b int) (q, r int) {
	return a / b, a % b
}

func apply(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("Calling function %s with args"+"(%d,%d)\n", opName, a, b)
	return op(a, b)
}

func sum(numbers ...int) int {
	s := 0
	for i := range numbers {
		s += numbers[i]
	}
	return s
}

func main() {
	fmt.Println(div(13, 3))
	fmt.Println(calculate(3, 4, "e"))
	fmt.Println(sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 4, 15, 16, 17, 18, 19, 20))
}
