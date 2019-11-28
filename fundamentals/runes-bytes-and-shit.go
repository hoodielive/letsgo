package main

import "fmt"

func main() {
	my_byte := byte('b')
	my_rune := '~'
	fmt.Printf("%c = %d   %c = %U\n",
		my_byte, my_byte, my_rune, my_rune)
}
