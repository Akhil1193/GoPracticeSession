package main

import (
	"fmt"
	"math"
)

func main() {

	const inflationRate = 3.5
	var investmentAmount float64
	var returnRate float64
	years := 10.0

	outputStatment("investment amount: ")
	fmt.Scan(&investmentAmount)

	outputStatment("returnRate: ")
	fmt.Scan(&returnRate)

	futureValue := float64(investmentAmount) * math.Pow(1+returnRate/100, years)
	futureReturnValue := futureValue / math.Pow(1+inflationRate/100, years)

	fv := fmt.Sprintf("future value is %.f\n", futureValue)
	frv := fmt.Sprintf("future return value is %.2f", futureReturnValue)

	fmt.Println(fv, frv)

	// fmt.Printf("future value is %v\n future return value is %v",futureValue,futureReturnValue)
	//fmt.Printf("future return valueis %v",futureReturnValue)

	//fmt.Printf("future value is %.f\n future return value is %.2f",futureValue,futureReturnValue)

}

func outputStatment(text string) {
	fmt.Println(text)
}
