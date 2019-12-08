package main

import "fmt"

func main() {
	// Define a map.
	moons := make(map[string]string)

	// Assign.
	moons["Earth"] = "Moon"
	moons["Jupiter"] = "Europa"
	moons["Saturn"] = "Titan"
	fmt.Println(moons)

	// Delete.
	delete(moons, "Saturn")
	fmt.Println(moons)
}
