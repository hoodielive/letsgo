package main

import "fmt"

func zero(x int) {
	x = 0
}

func zeroPointer(xPtr *int) {
	*xPtr = 0
}

func main() {
	x := 5
	zero(x)
	zeroPointer(&x)
	fmt.Println(x)
}
