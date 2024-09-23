package main

import (
	"bufio"
	"fmt"
	"linearstats/mathskills"
	"os"
	"strconv"
	"strings"
)




func main(){args := os.Args[1:]
	if len(args) != 1 {
		if len(args) > 1 {
			fmt.Println("Too many arguments. Please provide one file name.")
		} else {
			fmt.Println("Less arguments. Please enter the data file name.")
		}
		os.Exit(1)
	}

	input := args[0]
	file, err := os.Open(input)
	if err != nil {
		fmt.Println("Error opening the file:", err)
		return
	}
	defer file.Close()

	var dataSlice []float64
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		
		num, err := strconv.ParseFloat(line, 64)
		if err != nil {
			fmt.Println("Skipping invalid number:", line)
			continue
		}
		// Append the valid numbers to the slice
		dataSlice = append(dataSlice, num)
	}

	// Check for scanning errors
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading the file:", err)
		return
	}

	if len(dataSlice) > 1 {
		coefK, coefB := mathskills.LinearRegLine(dataSlice)
		fmt.Print("Linear Regression Line: y = ")
		fmt.Printf("%.6f", coefK)
		fmt.Print("x + ")
		fmt.Printf("%.6f", coefB)
		fmt.Println()

		pearson := mathskills.PearsonCoef(dataSlice)
		fmt.Print("Pearson Correlation Coefficient: ")
		fmt.Printf("%.10f", pearson)
		fmt.Println()
	}
}