package main

import "fmt"

//Naming in package
var aa = 3
var ss = "kkk"
var (
	cc = true
	dd = 4
)

//bb:=true    is not correct

func variableZeroValue() {
	var a int
	var s string
	fmt.Printf("%d %q\n", a, s)
}

func variableInitialValue() {
	var a, b int = 3, 4
	var s string = "abc"
	fmt.Println(a, b, s)
}

func variableTypeDeduction() {
	var a, b, c, s = 3, 4, true, "def"
	fmt.Println(a, b, c, s)
	fmt.Println(aa, cc, dd)
}

func variableShorter() {
	a, b, c, s := 3, 4, true, "def"
	b = 5
	a = 4
	fmt.Println(a, b, c, s)
}
func main() {
	fmt.Println("Hello World!")
	variableZeroValue()
	variableInitialValue()
	variableTypeDeduction()
	variableShorter()
	fmt.Println(aa, cc, dd, ss)
}
