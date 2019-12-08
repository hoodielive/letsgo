package main

import "fmt"

// The following struct represents a type with different fields.

type example struct {
	// 8 bytes | padding from flag
	flag    bool
	counter int16
	pi      float32
}

var e1 struct {
	flag    bool
	counter int16
	pi      float32
}

func main() {
	// A literal construction.
	e2 := example{
		flag:    true,
		counter: 10,
		pi:      3.141592,
	}

	e3 := struct {
		flag    bool
		counter int16
		pi      float32
	}{
		flag:    true,
		counter: 10,
		pi:      3.141592,
	}

	// Lets display the field values.
	fmt.Println("Flag", e2.flag)
	fmt.Println("Counter", e2.counter)
	fmt.Println("pi", e2.pi)

	fmt.Println("e3 Flag", e3.flag)
}
