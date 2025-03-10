package main

import (
	"fmt"
	"os"
)

var f = fmt.Println

func main() {
	accountBalance := 1000.0
	for {
		fmt.Println("Welcome to the Bank of Go!")
		fmt.Println("What would you like to do?")
		fmt.Println("1. Check Balance")
		fmt.Println("2. Deposit money")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit")

		var choice int
		fmt.Print("Your Choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			f("Your balance is", accountBalance)
		case 2:
			f("How much would you like to deposit?")
			var depositAmount float64
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				f("Invalid deposit amount!")
				continue
			}

			accountBalance += depositAmount
			writingBalanceToFile(accountBalance)
			f("Balance after deposit is", accountBalance)
		case 3:
			f("How much would you like to withdraw?")
			var withdrawAmount float64
			fmt.Scan(&withdrawAmount)

			if withdrawAmount <= 0 {
				f("Invalid withdrawal amount!")
				continue
			}

			if withdrawAmount > accountBalance {
				f("Not enough money!")
				return
			}
			accountBalance -= withdrawAmount
			writingBalanceToFile(accountBalance)
			f("Balance after withdrawal is", accountBalance)
		default:
			f("Bye!")
			return
		}
	}
}

func writingBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile("balance.txt", []byte(balanceText), 0644)
}
