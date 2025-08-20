package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func main() {

	operation := inputOperation("input operation please (AVG, SUM, MED):")
	transactionsString := inputTransactions("input transactions plz. Transaction must be input with comma (example: 2,30,10,10,-1 etc):")
	transactions := splitBySeparator(transactionsString, ",")
	result := actionWithTransations(transactions, operation)
	fmt.Println(result)

}

func actionWithTransations(tr []int, action string) int {
	processedAction := strings.ToUpper(action)

	switch processedAction {
	case "SUM":
		var result int
		for _, v := range tr {
			result += v
		}
		return result
	case "AVG":
		var result int
		for _, v := range tr {
			result += v
		}
		return result / len(tr)
	default:
		slices.Sort(tr)
		if len(tr) <= 1 {
			return tr[0]
		}
		if len(tr)/2 == 0 {
			min := len(tr) / 2
			max := len(tr)/2 - 1
			return (tr[min] + tr[max]) / 2
		}
		return tr[len(tr)/2]
	}
}

func splitBySeparator(text, sep string) []int {
	transactions := strings.Split(text, sep)
	result := make([]int, len(transactions))

	for i, v := range transactions {
		d, err := strconv.Atoi(v)
		if err != nil {
			fmt.Println(err)
		}
		result[i] = d
	}
	return result
}

func inputTransactions(text string) (transactions string) {
	fmt.Println(text)
	fmt.Scan(&transactions)
	return
}

func inputOperation(text string) (operation string) {
	fmt.Println(text)
	for {
		fmt.Scan(&operation)

		if operation != "SUM" && operation != "MED" && operation != "AVG" && operation != "sum" && operation != "med" && operation != "avg" {
			fmt.Println(`Repeat plz operation`)
			continue
		}
		break
	}
	return
}
