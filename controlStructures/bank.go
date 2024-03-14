package main

import "fmt"

func main() {
	accoutBalance := 1000.0

	for {
		fmt.Println("Welcome to Go Bank!")
		fmt.Println("What do you want to do?")
		fmt.Println("1.Check Balance")
		fmt.Println("2.Deposit Amount")
		fmt.Println("3.Withdraw Amount")
		fmt.Println("4.Exit")

		var Choice int

		fmt.Print("Enter your choice:")
		fmt.Scan(&Choice)

		if Choice == 1 {
			fmt.Println("Your account balance is:", accoutBalance)
		} else if Choice == 2 {
			fmt.Println("Your Deposit Amount:")
			var addAmount float64
			fmt.Scan(&addAmount)
			if addAmount <= 0 {
				fmt.Println("Invalid Amount! Amount should be greater than Zero")
				continue
			}
			accoutBalance += addAmount
			fmt.Println("Your total balance is:", accoutBalance)
		} else if Choice == 3 {
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

			fmt.Println("Your total balance is:", accoutBalance)
		} else {
			fmt.Println("Good Bye!")
			//return
			break
		}
	}
	fmt.Println("Thankyou for banking with us")
}