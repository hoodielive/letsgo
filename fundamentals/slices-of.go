package main

import "fmt"

func main() {
	alphabet := []string{"a", "b", "c", "d", "e"}
	fmt.Println(alphabet)

	// make (type, len, cap)
	s := make([]int, 5, 10)
	fmt.Println("len=%d cap=%d %v\n", len(s), cap(s), s)
}
