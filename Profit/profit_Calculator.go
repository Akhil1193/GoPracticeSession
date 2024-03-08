package main

import "fmt"

func main() {
	var revenue float64
	var expenses float64
	var taxRate float64

	textOuput("revenue:")
	fmt.Scan(&revenue)

	textOuput("expenses:")
	fmt.Scan(&expenses)

	textOuput("taxRate: ")
	fmt.Scan(&taxRate)

	// ebt := (revenue - expenses) * (1 - taxRate)
	// incomeTaxExpense := ebt * taxRate
	// eat := ebt - incomeTaxExpense
	// ratio := ebt / eat
	Ebt,IncomeTaxExpense,Eat,Ratio:=calculation(revenue,expenses,taxRate)

	fmt.Println(Ebt)
	fmt.Println(IncomeTaxExpense)
	fmt.Println(Eat)
	fmt.Println(Ratio)
}

func textOuput(text string) string {
	fmt.Print(text)
	return text
}

func calculation(revenue,expenses, taxRate float64) (ebt,incomeTaxExpense,eat,ratio float64){
	ebt = (revenue - expenses) * (1 - taxRate)
	incomeTaxExpense = ebt * taxRate
	eat = ebt - incomeTaxExpense
	ratio = ebt / eat
    return ebt,incomeTaxExpense,eat,ratio
}
