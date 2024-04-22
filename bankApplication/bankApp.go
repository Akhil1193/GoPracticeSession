package main

import (
	"fmt"
	"GoPracticeSession/bankApplication/fileops"
	"github.com/Pallinder/go-randomdata"
)

const accountBalance = "balance.txt"

func main() {

	accountBalanc, err := fileops.GetFloatFromFile(accountBalance)

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("--------")
		//panic(err)
	}

	fmt.Println("Welcome to HDFC bank!")
	fmt.Println(randomdata.PhoneNumber())

	for {

		readOptions()
		var choice int
		

		fmt.Print("Your choice:")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Your account balance:", accountBalanc)
		case 2:
			fmt.Print("Enter your amount to deposit:")
			var depositAmount float64
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Entered invalid amount")
				continue
			}
			accountBalanc += depositAmount
			fileops.WriteFloatToFile(accountBalanc, accountBalance)
			fmt.Println("Updated balance:", accountBalanc)
		case 3:
			fmt.Println("Enter your amount to withdraw:")
			var withdrawAmount float64
			fmt.Scan(&withdrawAmount)

			if withdrawAmount <= 0 {
				fmt.Println("Enter correct Amount!")
				continue
			}

			if withdrawAmount > accountBalanc {
				fmt.Println("Amount is greater than Account Balance!")
				continue
			}
			accountBalanc -= withdrawAmount
			fileops.WriteFloatToFile(accountBalanc, accountBalance)
			fmt.Println("Updated balance:", accountBalanc)
		default:
			fmt.Println("Goodbye!")
			fmt.Println("Thankyou for banking with us!")
			return
		}

	}
}
