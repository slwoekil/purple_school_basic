package main

import "fmt"

func main() {
	const USDtoRUB = 80.70
	const USDtoEUR = 0.86
	EURtoRUB := USDtoRUB / USDtoEUR
	fmt.Printf("%.2f", EURtoRUB)
}
func getCurrencies() {
	var currencyToChange string
	var currencyToGet string
	var countCurrency float64

	fmt.Println("Какую валюту вы хотите обменять?")
	fmt.Scan(currencyToChange)
	fmt.Println("Какую валюту вы хотите получить?")
	fmt.Scan(currencyToGet)
	fmt.Println("Какое количество валюты вы хотите обменять?")
	fmt.Scan(countCurrency)
}

func calculateChangedCurrency(currencyToChange string, currencyToGet string, countCurrency float64) {

}
