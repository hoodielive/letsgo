package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("One step at a time.")

	beyondOne()
}

func beyondOne() {
	var x int
	x = 3
	y := 4
	sum, prod := learnMultiple(x, y)
	fmt.Println("sum: ", sum, "prod: ", prod)
	learnTypes()
}

func learnMultiple(x, y int) (sum, prod int) {
	return x + y, x * y
}

func learnTypes() {
	str := "Learn Go!"
	s2 := `A "raw" string literal 
				can include line breaks.`
	g := '∑'
	f := 3.14195
	c := 3 + 4i
	var u uint = 7
	var pi float32 = 22. / 7
	n := byte('\n')
	var a4 [4]int
	a5 := [...]int{3, 1, 5, 10, 100}
	s3 := []int{4, 5, 9}
	s4 := make([]int, 4)
	var d2 [][]float64
	bs := []byte("a slice")
	s := []int{1, 2, 3}
	s = append(s, 4, 5, 6)
	fmt.Println(s)
	p, q := learnMemory()
	fmt.Println(*p, *q)

	m := map[string]int{"three": 3, "four": 4}
	m["one"] = 1

	_, _, _, _, _, _, _, _, _, _ = str, s2, g, f, u, pi, n, a5, s4, bs

	file, _ := os.Create("output.txt")
	fmt.Fprint(file, "This is how you write to a file, by the way")
	file.Close()

	fmt.Println(s, c, a4, s3, d2, m)

	learnFlowControl()
}

func learnNamedReturns(x, y int) (z int) {
	z = x * y
	return
}

func learnMemory() (p, q *int) {
	// named return values p and q have type pointer to int
	p = new(int)
	s := make([]int, 20)
	s[3] = 7
	r := -2
	return &s[3], &r
}

func expensiveComputation() float64 {
	return m.Exp(10)
}

func learnFlowControl() {
	if true {
		fmt.Println("Told ya.. ")
	}

	if false {
		 // Pout.
	}
	else {
		// Gloat.
	}

	x := 42.0
	switch x {
	case 0:
	case 1:
	case 42:
		// Cases don't fall through
		/*
		this is a fallthrough keyword however, see
		*/
	case 43:
		// Unreached
	default:
		// Default case.
	}

	for x: 0; x < 3; x++ {
		fmt.Println("Iteration..", x)
	}

	for {
		break // Just kidding.
		continue // Unreached.
	}

	for key, value := range map[string]int{"one": 1, "two": 2, "three": 3} {
		fmt.Printf("key=%s, value=%d\n", key, value)
	}

	for _, name := range []string{"Osa", "Okan", "Okanran"} {
		fmt.Print("Hello, %s\n", name)
	}

	if y := expensiveComputation(); y > x {
		x = y
	}

	xBig := func() bool {
		return x > 10000
	}

	x = 99999
	fmt.Println("xBig:", xBig())
}