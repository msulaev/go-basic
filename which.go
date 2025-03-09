package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide an argument!")
		os.Exit(1)
	}
	file := arguments[1]
	path := os.Getenv("PATH")
	pathSplit := filepath.SplitList(path)
	found := false
	for _, directory := range pathSplit {
		fullPath := filepath.Join(directory, file)
		fileInfo, err := os.Stat(fullPath)
		if err != nil {
			continue
		}

		mode := fileInfo.Mode()
		if !mode.IsRegular() {
			continue
		}
		if mode&0111 != 0 {
			fmt.Println(fullPath)
			os.Exit(0)
		}
	}
	if !found {
		fmt.Println("File not found")
		os.Exit(1)
	}
	os.Exit(0)
}
