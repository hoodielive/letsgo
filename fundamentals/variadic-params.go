package main

import "fmt"

func VariadicParams(myStrings ...interface{}) {
	// iterate each value of the variadic.
	for _, param := range myStrings {
		fmt.Println("param:", param)
	}

	fmt.Println("params:", fmt.Sprintln(myStrings...))

	learnErrorHandling()
}

func learnErrorHandling() {
	m := map[int]string{3: "three", 4: "four"}

	if x, ok := m[1]; !ok {
		fmt.Println("No other there...")
	}
	else {
		fmt.Print(x)
	}

	if _, err := strconv.Atoi("non-int"); err != nil {
		fmt.Println(err)
	}
  }
}