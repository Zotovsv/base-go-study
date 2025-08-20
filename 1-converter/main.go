package main

import (
	"fmt"
)

const (
	USDToEUR = 0.75
	USDToRUB = 82.30
)

func main() {
	currentCurrency := inputCurrency("Input current currency: USD,EUR, RUB", "")
	amount := inputSum("Input amount current currency")
	targetCurrency := inputCurrency("Input target currency", currentCurrency)

	result := calculate(amount, currentCurrency, targetCurrency)
	fmt.Println(result)
}

func calculate(amount float64, current, target string) float64 {
	var toUSD float64
	switch current {
	case "USD":
	case "usd":
		toUSD = 1
	case "EUR":
	case "eur":
		toUSD = 1 / USDToEUR
	case "RUB":
	case "rub":
		toUSD = 1 / USDToRUB
	default:
		return 0
	}

	usd := amount / toUSD // конвертируем текущую валюту в USD

	switch target {
	case "USD":
	case "usd":
		return usd
	case "EUR":
	case "eur":
		return usd * USDToEUR
	case "RUB":
	case "rub":
		return usd * USDToRUB
	default:
		return 0
	}
	return usd
}

func inputCurrency(text, currentCurrency string) (currency string) {
	fmt.Println(text)
	for {
		fmt.Scan(&currency)
		if currency == currentCurrency {
			fmt.Printf("currentCurrent and currency must be different. Current %v, Target %v\n", currentCurrency, currency)
			continue
		}
		if currency != "EUR" && currency != "RUB" && currency != "USD" && currency != "eur" && currency != "rub" && currency != "usd" {
			fmt.Println(`Repeat plz currency`)
			continue
		}
		break
	}
	return
}

func inputSum(text string) float64 {
	var sum float64
	fmt.Println(text)
	for {
		fmt.Scan(&sum)
		if sum <= 0 {
			fmt.Println(`Repeat plz`)
			continue
		}
		break
	}
	return sum
}
