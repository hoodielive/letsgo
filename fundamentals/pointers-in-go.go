package main

import "fmt"

func main() {
	x, n := 32, 909
	z := &x
	fmt.Println(*z)
	*z = 82
	fmt.Println(x)

	z = &n
	*z /= 37
	fmt.Println(n)
}
