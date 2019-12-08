package main

import "fmt"

func main() {
	atr := make(map[string]string)

	atr["Haiti"] = "Vodou"
	atr["Ghana"] = "Fa"
	atr["Yoruba"] = "Ifa"
	atr["Congo"] = "Palo Mayombe"

	fmt.Println(atr)
	atr_online := map[string]string{"Haiti": "Vodou", "Yoruba": "Ifa", "Congo": "Palo Mayombe"}
	fmt.Println(atr_online)
}
