package main

import "fmt"

func main() {
	// rule of thumb: use var when variable is in a zero state and shorthand
	// when its not.

	var a int
	var b string
	var c float64
	var d bool

	fmt.Printf("var a int \t %T [%v]\n", a, a)
	fmt.Printf("var b string \t %T [%v]\n", b, b)
	fmt.Printf("var c float64 \t %T [%v]\n", c, c)
	fmt.Printf("var d bool \t %T [%v]\n", d, d)

	aa := 10
	bb := "hello"
	cc := 3.14159
	dd := true

	fmt.Printf("aa is \t %T [%v]\n", aa, aa)
	fmt.Printf("bb is \t %T [%v]\n", bb, bb)
	fmt.Printf("cc is \t %T [%v]\n", cc, cc)
	fmt.Printf("dd is \t %T [%v]\n", dd, dd)

	// specify type and perform a conversion || conversion over casting
	aaa := int32(10)
	fmt.Printf("aaa is \t %T [%v]\n", aaa, aaa)
}
