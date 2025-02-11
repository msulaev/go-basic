package main

import (
	"fmt"
)

func main() {
	var revenue, expenses, taxRate float64
	fmt.Print("Enter revenue:")
	fmt.Scan(&revenue)
	fmt.Print("Enter expenses:")
	fmt.Scan(&expenses)
	fmt.Print("Enter tax rate:")
	fmt.Scan(&taxRate)
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / revenue
	fmt.Println("EBT: ", ebt)
	fmt.Println("Profit: ", profit)
	fmt.Println("Ratio: ", ratio)
}
