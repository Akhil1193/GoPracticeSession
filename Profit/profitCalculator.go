package main

import (
	"errors"
	"fmt"
	"os"
)

func writeDataToFile(ebt, profit, ratio float64) {

	fileName := "dataText.txt"
	dataText := fmt.Sprintf("EBT:%.1f\nPROFIT:%.1f\nRATIO:%.3f\n",ebt, profit, ratio)
	os.WriteFile(fileName, []byte(dataText), 0644)
}

func main() {
	revenue, err1 := getUserInput("Revenue: ")

	if err1 != nil{
		return
	}
	expenses, err2 := getUserInput("Expenses: ")

	if err2 !=nil{
		panic(err2)
	}

	taxRate, err3 := getUserInput("Tax Rate: ")
    if err3 !=nil{
		panic(err3)
	}

	ebt, profit, ratio := calculateFinancials(revenue, expenses, taxRate)

	writeDataToFile(ebt, profit, ratio)

	fmt.Printf("EBT is %.1f\n", ebt)
	fmt.Printf("Profit is %.1f\n", profit)
	fmt.Printf("Ratio is %.3f\n", ratio)
}

func calculateFinancials(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit
	return ebt, profit, ratio
}

func getUserInput(infoText string) (float64, error) {
	var userInput float64

	fmt.Print(infoText)
	fmt.Scan(&userInput)

	if userInput <= 0 {
		return 0, errors.New("value must be positive number")
	}

	return userInput, nil
}
