package main

import "fmt"

func main() {
	const USDtoRUB = 80.70
	const USDtoEUR = 0.86
	EURtoRUB := USDtoRUB / USDtoEUR
	fmt.Printf("%.2f", EURtoRUB)
}
