package main

import (
	"fmt"
)

const USDtoRUB = 80.70
const USDtoEUR = 0.86
const EURtoRUB = USDtoRUB / USDtoEUR

func main() {
	currencyForChange := getCurrencyForChange()
	countCurrency := getCountOfCurrency()
	receivedCurrency := getReceivedCurrency()
	calculateChangedCurrency(currencyForChange, receivedCurrency, countCurrency)
}

func calculateChangedCurrency(currencyForChange string, receivedCurrency string, countCurrency float64) {
	var exchange string

	switch {
	case currencyForChange == "EUR" && receivedCurrency == "USD":
		exchange = fmt.Sprintf("Сумма к получению: %.2f", countCurrency/USDtoEUR)
		fmt.Println(exchange, receivedCurrency)
	case currencyForChange == "EUR" && receivedCurrency == "RUB":
		exchange = fmt.Sprintf("Сумма к получению: %.2f", EURtoRUB*countCurrency)
		fmt.Println(exchange, receivedCurrency)
	case currencyForChange == "EUR" && receivedCurrency == "EUR":
		fmt.Println("Вы выбрали две одинаковые валюты")
	case currencyForChange == "USD" && receivedCurrency == "EUR":
		exchange = fmt.Sprintf("Сумма к получению: %.2f", USDtoEUR*countCurrency)
		fmt.Println(exchange, receivedCurrency)
	case currencyForChange == "USD" && receivedCurrency == "RUB":
		exchange = fmt.Sprintf("Сумма к получению: %.2f", USDtoRUB*countCurrency)
		fmt.Println(exchange, receivedCurrency)
	case currencyForChange == "USD" && receivedCurrency == "USD":
		fmt.Println("Вы выбрали две одинаковые валюты")
	case currencyForChange == "RUB" && receivedCurrency == "EUR":
		exchange = fmt.Sprintf("Сумма к получению: %.2f", countCurrency/EURtoRUB)
		fmt.Println(exchange, receivedCurrency)
	case currencyForChange == "RUB" && receivedCurrency == "USD":
		exchange = fmt.Sprintf("Сумма к получению: %.2f", countCurrency/USDtoRUB)
		fmt.Println(exchange, receivedCurrency)
	case currencyForChange == "RUB" && receivedCurrency == "RUB":
		fmt.Println("Вы выбрали две одинаковые валюты")
	}
}

func getCurrencyForChange() string {
	var currencyForChange string

	for {
		fmt.Println("Какую валюту вы хотите обменять? (USD, EUR, RUB)")
		fmt.Scan(&currencyForChange)

		if currencyForChange == "USD" || currencyForChange == "EUR" || currencyForChange == "RUB" {
			break
		}

		fmt.Println("Валюта не найдена")
	}

	return currencyForChange
}

func getCountOfCurrency() float64 {
	var countCurrency float64

	for {
		fmt.Println("Какое количество валюты вы хотите обменять?")

		_, err := fmt.Scan(&countCurrency)

		if err != nil {
			fmt.Println("Вы ввели некорректное значение. Для продолжения введите число")
			var discard string
			fmt.Scan(&discard)
			continue
		}

		if countCurrency > 0 {
			break
		}

		fmt.Println("Сумма для обмена должна быть больше 0")
	}

	return countCurrency
}

func getReceivedCurrency() string {
	var receivedCurrency string

	for {
		fmt.Println("Какую валюту вы хотите получить? (USD, EUR, RUB)")
		fmt.Scan(&receivedCurrency)

		if receivedCurrency == "USD" || receivedCurrency == "EUR" || receivedCurrency == "RUB" {
			break
		}

		fmt.Println("Валюта не найдена")
	}

	return receivedCurrency
}
