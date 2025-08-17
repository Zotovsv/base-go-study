package main

import "fmt"

func main() {
	const USDToEURO = 0.93
	const USDToRUB = 92.50

	const EUROS = 100.0
	const usd = EUROS / USDToEURO
	const rubs = usd * USDToRUB
	fmt.Print(rubs)

}
