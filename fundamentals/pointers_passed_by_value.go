package main

import "fmt"

func main() {
	count := 10
	fmt.Println("Count:\tValue Of[", count, "]\tAddr Of[", &count, "]")
	increment(count)
	fmt.Println("Count:\tValue Of[", count, "]\tAddr Of[", &count, "]")
}
