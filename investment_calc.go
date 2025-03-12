package main

import (
	"errors"
	"fmt"
	"os"
)

var file = "result.txt"

func main() {
	revenue, err := getUserInput("Enter revenue: ")
	showErrorMessage(err)

	expenses, err := getUserInput("Enter expenses: ")
	showErrorMessage(err)

	taxRate, err := getUserInput("Enter tax rate: ")
	showErrorMessage(err)

	profit := calculateProfit(revenue, expenses, taxRate)
	err = writeToFile(profit)
	if err != nil {
		showErrorMessage(err)
		return
	}

	formattedFV := fmt.Sprintf("Future Value: %.1f\n", profit)
	fmt.Println(formattedFV)
}

func calculateProfit(revenue, expenses, taxRate float64) float64 {
	ebt := revenue - expenses
	return ebt * (1 - taxRate/100)
}

func getUserInput(prompt string) (float64, error) {
	fmt.Print(prompt)
	var input float64
	_, err := fmt.Scan(&input)
	if err != nil || input <= 0 {
		return 0, errors.New("invalid input")
	}
	return input, nil
}

func writeToFile(result float64) error {
	err := os.WriteFile(file, []byte(fmt.Sprintf("%f", result)), 0644)
	if err != nil {
		return errors.New("failed to write to file")
	}
	return nil
}

func showErrorMessage(err error) {
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
}
