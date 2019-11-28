package main

import "fmt"

func main() {
	s := []byte("abc")
	// print literal byte
	fmt.Println(s)

	// append a byte
	s = append(s, byte('d'))

	// print string representation of the byte
	fmt.Println(string(s))

	// length of the byte slice
	fmt.Println(len(s))

	// 1st byte
	fmt.Println(s[0])
}
