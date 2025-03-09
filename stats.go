package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	argumetns := os.Args
	if len(argumetns) == 1 {
		fmt.Println("Please provide one integer argument!")
		return
	}
	var min, max float64
	var initialized = 0

	nValues := 0
	var sum float64
	for i := 1; i < len(argumetns); i++ {
		n, err := strconv.ParseFloat(argumetns[i], 64)
		if err != nil {
			continue
		}
		nValues = nValues + 1
		sum = sum + n
		if initialized == 0 {
			min = n
			max = n
			initialized = 1
		}
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}
	fmt.Println("Number is ", nValues)
	fmt.Println("Min is ", min)
	fmt.Println("Max is ", max)
	//mean value

	if nValues == 0 {
		return
	}
	meanValue := sum / float64(nValues)
	fmt.Printf("Mean is %.5f\n", meanValue)

	var squared float64
	for i := 1; i < len(argumetns); i++ {
		n, err := strconv.ParseFloat(argumetns[i], 64)
		if err != nil {
			continue
		}
		squared = squared + math.Pow((n-meanValue), 2)
	}
	stdv := math.Sqrt(squared / float64(nValues))
	fmt.Printf("Standard deviation is %.5f\n", stdv)
}
