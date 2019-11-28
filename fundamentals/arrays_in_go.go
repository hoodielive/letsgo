package main

import "fmt"

func main() {
	var fruits [2]string
	fruits[0] = "pomegrante"
	fruits[1] = "rambutan"
	fmt.Println(fruits)

	smoothies := [2]string{"avocado", "peanut butter"}
	fmt.Println(smoothies)

	ids := []int{1, 3, 5, 7, 9}
	fmt.Println(ids)
}
