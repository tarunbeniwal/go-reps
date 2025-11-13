package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	// take input from user

	scanner := bufio.NewScanner(os.Stdin)

	var num1, num2 float64
	var err error
	for {
		// First input
		for {

			fmt.Printf("Enter a number: ")
			scanner.Scan()
			input1 := scanner.Text()

			num1, err = strconv.ParseFloat(input1, 64)
			if err != nil {
				fmt.Printf("Invalid input. Please enter a valid number. \n")
				continue
			}
			break
		}

		// Next input
		for {

			fmt.Printf("Enter another number: ")
			scanner.Scan()
			input2 := scanner.Text()

			num2, err = strconv.ParseFloat(input2, 64)
			if err != nil {
				fmt.Printf("Invalid input. Please enter a valid number. \n")
				continue
			}
			break
		}

		// operation
		for {
			fmt.Printf("Enter an operator (+, -, *, /): ")
			scanner.Scan()
			operator := scanner.Text()

			switch operator {

			case "+":
				fmt.Printf("Sum of numbers are: %.2f \n", num1+num2)

			case "-":
				fmt.Printf("Difference of numbers are: %.2f \n", num1-num2)

			case "*":
				fmt.Printf("Product of numbers are: %.2f \n", num1*num2)

			case "/":
				if num2 == 0 {
					fmt.Printf("Divisor can't be zero \n")
				} else {
					fmt.Printf("Division of numbers are: %.2f \n", num1/num2)

				}

			default:
				fmt.Printf("Invalid operator. Please use one of +, -, *, /. \n")
				continue
			}
			break
		}
		// ask if user wants to calculate again
		fmt.Printf("Do you want to calulate again? (y/n): ")
		scanner.Scan()
		again := strings.ToLower(scanner.Text())
		if again != "y" {
			fmt.Printf("Goodbye! \n")
			break
		}
	}
}
