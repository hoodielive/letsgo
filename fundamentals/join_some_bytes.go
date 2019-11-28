package main

import (
	"bytes"
	"fmt"
)

func main() {
	b1 := byte('a')
	b2 := []byte("A")
	b3 := []byte{'a', 'b', 'c'}

	fmt.Printf("b1 = %c\n", b1)
	fmt.Printf("b2 = %c\n", b2)
	fmt.Printf("b3 = %c\n", b3)

	s1 := []byte("hello")
	s2 := []byte("Aeon")
	s3 := [][]byte{s1, s2}
	s4 := bytes.Join(s3, []byte(","))
	s5 := []byte{}
	s6 := bytes.Join(s3, []byte("--"))
	s7 := [][]byte{[]byte("foo"), []byte("bar"), []byte("baz")}

	fmt.Printf("s1 = %s\n", s1)
	fmt.Printf("s2 = %s\n", s2)
	fmt.Printf("s3 = %s\n", s3)
	fmt.Printf("s4 = %s\n", s4)
	fmt.Printf("s5 = %s\n", s5)
	fmt.Printf("s6 = %s\n", s6)
	fmt.Printf("%s\n", bytes.Join(s7, []byte(", ")))
}
