package main

import (
	"GoPracticeSession/controlStructures/fileops"
	"github.com/Pallinder/go-randomdata"
	"fmt"
)

const accountBalanceFile = "balance.txt"

func main() {
	accoutBalance, err := fileops.GetFloatFromFile(accountBalanceFile)

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("---------")
	}

	fmt.Println("Welcome to HDFC Bank!")
	fmt.Println("Email us your concerns at",randomdata.Email())
	fmt.Println("Reach us 24/7 at",randomdata.PhoneNumber())
	fmt.Println(randomdata.Address())

	for {
		fileops.PrintOptions()

		var Choice int

		fmt.Print("Enter your choice:")
		fmt.Scan(&Choice)

		switch Choice {
		case 1:
			fmt.Println("Your account balance is:", accoutBalance)
		case 2:
			fmt.Println("Your Deposit Amount:")
			var addAmount float64
			fmt.Scan(&addAmount)
			if addAmount <= 0 {
				fmt.Println("Invalid Amount! Amount should be greater than Zero")
				continue
			}
			accoutBalance += addAmount
			fmt.Println("Your total balance is:", accoutBalance)
			fileops.WriteFloatToFile(accoutBalance, accountBalanceFile)
		case 3:
			fmt.Println("Withdraw Amount:")
			var WithdrawAmount float64
			fmt.Scan(&WithdrawAmount)
			if WithdrawAmount <= 0 {
				fmt.Println("Invalid Amount! Amount should be greater than Zero")
				continue
			}
			if WithdrawAmount > accoutBalance {
				fmt.Println("Insufficient Amount in your Account!")
				continue
			}
			accoutBalance -= WithdrawAmount
			fileops.WriteFloatToFile(accoutBalance, accountBalanceFile)

			fmt.Println("Your total balance is:", accoutBalance)
		default:
			fmt.Println("Good Bye!")
			fmt.Println("Thankyou for banking with us")
			return

			//break
		}

	}
}
