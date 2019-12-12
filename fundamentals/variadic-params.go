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
	learnConcurrency()
  }

  func inc(i int , c chan int) {
	  c <- i + 1
  }

  func learnConcurrency() {
	  c := make(chan int)

	  go inc(0, c)
	  go inc(10, c)
	  go inc(-805, c)

	  fmt.Println(<-c, <-c, <-c)

	  cs := make(chan string)
	  ccs := make(chan chan string)
	  go func() { c <- 84 }() 
	  go func() { c <- 'wordy' }()

	  select {
	  case i := <-c:
		fmt.Printf("It's a %T", i)
	  case <-cs:
		fmt.Println("It's a string")
	  case <-ccs:
		fmt.Println("Didnt happen..")
	  }

	  learnWebProgramming()
  }
}