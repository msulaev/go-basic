package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

var f = fmt.Println

const balanceFile = "balance.txt"

func main() {
	accountBalance, err := getBalanceFromFile()
	if err != nil {
		f("Error getting balance from file")
		f(err)
		panic("error getting balance from file")
	}
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
	os.WriteFile(balanceFile, []byte(balanceText), 0644)
}

func getBalanceFromFile() (float64, error) {
	data, err := os.ReadFile(balanceFile)
	if err != nil {
		return 1000, errors.New("error reading balance file")
	}
	balance, err := strconv.ParseFloat(string(data), 64)
	if err != nil {
		return 1000, errors.New("error converting balance to float")
	}
	return balance, nil
}
