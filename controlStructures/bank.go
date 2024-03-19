package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const fileName = "balance.txt"

func readBalanceFromFile() (float64, error) {
	data, err := os.ReadFile(fileName)

	if err != nil {
		return 1000, errors.New("failed to read file")
	}
	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)

	if err != nil {
		return 1000, errors.New("failed to parse value")
	}
	return balance, nil
}

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(fileName, []byte(balanceText), 0644)
}

func main() {
	accoutBalance, err := readBalanceFromFile()

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("---------")
	}

	fmt.Println("Welcome to HDFC Bank!")

	for {
		fmt.Println("Welcome to Go Bank!")
		fmt.Println("What do you want to do?")
		fmt.Println("1.Check Balance")
		fmt.Println("2.Deposit Amount")
		fmt.Println("3.Withdraw Amount")
		fmt.Println("4. Exit")

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
			writeBalanceToFile(accoutBalance)
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
			writeBalanceToFile(accoutBalance)

			fmt.Println("Your total balance is:", accoutBalance)
		default:
			fmt.Println("Good Bye!")
			fmt.Println("Thankyou for banking with us")
			return

			//break
		}

		// 	// if Choice == 1 {
		// 	// 	fmt.Println("Your account balance is:", accoutBalance)
		// 	// } else if Choice == 2 {
		// 	// 	fmt.Println("Your Deposit Amount:")
		// 	// 	var addAmount float64
		// 	// 	fmt.Scan(&addAmount)
		// 	// 	if addAmount <= 0 {
		// 	// 		fmt.Println("Invalid Amount! Amount should be greater than Zero")
		// 	// 		continue
		// 	// 	}
		// 	// 	accoutBalance += addAmount
		// 	// 	fmt.Println("Your total balance is:", accoutBalance)
		// 	// } else if Choice == 3 {
		// 	// 	fmt.Println("Withdraw Amount:")
		// 	// 	var WithdrawAmount float64
		// 	// 	fmt.Scan(&WithdrawAmount)
		// 	// 	if WithdrawAmount <= 0 {
		// 	// 		fmt.Println("Invalid Amount! Amount should be greater than Zero")
		// 	// 		continue
		// 	// 	}
		// 	// 	if WithdrawAmount > accoutBalance {
		// 	// 		fmt.Println("Insufficient Amount in your Account!")
		// 	// 		continue
		// 	// 	}
		// 	// 	accoutBalance -= WithdrawAmount
		// 	// 	fmt.Println("Your total balance is:", accoutBalance)
		// 	// } else {
		// 	// 	fmt.Println("Good Bye!")
		// 	// 	//return
		// 	// 	break
		// 	}

	}
}
