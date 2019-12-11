package main

import "fmt"

func zero(x int) {
	x = 0
}

func zeroPointer(xPtr *int) {
	*xPtr = 98
}

func main() {
	x := 5
	zero(x)
	fmt.Println(x)
	zeroPointer(&x)
	fmt.Println(x)
}
