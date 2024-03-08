package main

import "fmt"

func main() {
	var revenue float64
	var expenses float64
	var taxRate float64

	fmt.Print("revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("expense: ")
	fmt.Scan(&expenses)

	fmt.Print("taxRate: ")
	fmt.Scan(&taxRate)

	ebt := (revenue - expenses) * (1 - taxRate)
	incomeTaxExpense := ebt * taxRate
	eat := ebt - incomeTaxExpense
	ratio := ebt / eat

	fmt.Println(ebt)
	fmt.Println(incomeTaxExpense)
	fmt.Println(eat)
	fmt.Println(ratio)
}
