package main

import "fmt"

func main() {
	oleg := map[string]string{"name": "Hood", "animal": "shark", "color": "black"}
	fmt.Println(oleg["animal"])

	for key, value := range oleg {
		fmt.Printf("%q is the key for the value %q\n", key, value)
	}
}
