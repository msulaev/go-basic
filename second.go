package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide numbers")
		return
	}
	var total, nInts, nFloats int
	invalid := make([]string, 0)
	for _, k := range arguments[1:] {
		_, err := strconv.Atoi(k)
		if err != nil {
			invalid = append(invalid, k)
			continue
		}
		total++
		nInts++
	}
	fmt.Println("Total: ", total, " Ints: ", nInts, " Floats: ", nFloats)
	if len(invalid) > 0 {
		fmt.Printf("Invalid inputs: %v\n", len(invalid))
		for _, s := range invalid {
			fmt.Println(s)
		}
	}
}
