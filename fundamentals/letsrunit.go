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
