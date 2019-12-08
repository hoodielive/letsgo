package main

import "fmt"

func main() {
	atr := make(map[string]string)

	atr["Haiti"] = "Vodou"
	atr["Ghana"] = "Fa"
	atr["Yoruba"] = "Ifa"
	atr["Congo"] = "Palo Mayombe"

	fmt.Println(atr)
}
