package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	var num uint64
	var order string
	var err error

	// taking the number input ✅
	for {
		fmt.Printf("Enter a number to generate its multiplication table: ")
		scanner.Scan()
		input := scanner.Text()
		num, err = strconv.ParseUint(input, 10, 64)
		if err != nil {
			fmt.Println("Invalid input. Please enter a valid number.")
			continue
		}
		break
	}

	// taking the order input ✅
	for {

		fmt.Printf("Ascending or Descending? (a/d): ")
		scanner.Scan()
		input1 := scanner.Text()
		order = strings.ToLower(strings.TrimSpace(input1))
		if order != "a" && order != "d" {
			fmt.Println("Invalid input. Please enter 'a' for ascending or 'd' for descending.")
			continue
		}
		break
	}

	fmt.Printf("Multiplcation table of %d in %s order: \n", num, order)

	start, end, step := 1, 10, 1
	if order == "d" {
		start, end, step = 10, 1, -1
	}
	for i := start; ; i += step {
		fmt.Printf("%d x %d = %d \n", num, i, num*uint64(i))
		if i == end {
			break
		}
	}

	fmt.Println("Thank you!, Would you like to generate another table? (y/n)")
	scanner.Scan()
	response := scanner.Text()
	if strings.ToLower(strings.TrimSpace(response)) == "y" {
		main()
	} else {
		fmt.Println("Goodbye!")
	}

}
