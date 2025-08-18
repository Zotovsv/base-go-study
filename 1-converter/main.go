package main

import "fmt"

func main() {
	amount, current, target := inputData()
	fmt.Println(amount, current, target)
	result := calculate(amount, current, target)
	fmt.Println(result)
}

func calculate(amount float64, current, target string) (result float64) {
	return
}
func inputData() (amount float64, current, target string) {
	fmt.Println(`Input amount money:`)
	fmt.Scan(&amount)
	fmt.Println(`Input type money:`)
	fmt.Scan(&current)
	fmt.Println(`Input target money:`)
	fmt.Scan(&target)
	return
}
