package main

import (
	"bufio"
	calculatorproxy "distributed-platforms/internal/distribution/proxy"
	shared "distributed-platforms/internal/shared"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func calculation(num1 int, num2 int, operator string, c calculatorproxy.CalculatorProxy) {
	var result int
	switch operator {
	case "+":
		result = c.Sum(num1, num2)
	case "-":
		result = c.Sub(num1, num2)
	case "*":
		result = c.Mul(num1, num2)
	case "/":
		result = c.Div(num1, num2)
	default:
		fmt.Println("Invalid operator. Use one of: +, -, *, /")
		return
	}
	// Display the result
	fmt.Println("Result:", result)

}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Welcome to the Calculator!")
	fmt.Println("Enter your calculation in the format: number1 operator number2 (e.g., 12 + 5)")
	fmt.Println("Type 'exit' to quit, 'extend_lease' to keep using calculator")

	ior := shared.IOR{Host: shared.LocalHost, Port: shared.DefaultPort}
	c := calculatorproxy.New(ior)

	for {
		// Prompt user for input
		fmt.Print("> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}

		// Trim whitespace and check for 'exit'
		input = strings.TrimSpace(input)
		if input == "exit" {
			fmt.Println("Goodbye!")
			break
		}

		if input == "extend_lease" {
			fmt.Println("Command to extend lease!")
		}

		if input == "lease_type_0" {
			fmt.Println("TIPO 0")
		}
		if input == "lease_type_1" {
			fmt.Println("TIPO 1")
		}
		if input == "lease_type_2" {
			fmt.Println("TIPO 2")
		}

		// Split the input
		parts := strings.Split(input, " ")
		if len(parts) != 3 {
			fmt.Println("Invalid input format. Use: number1 operator number2")
			continue
		}

		// Parse the numbers
		num1, err := strconv.ParseInt(parts[0], 10, 0)
		if err != nil {
			fmt.Println("Invalid number:", parts[0])
			continue
		}

		num2, err := strconv.ParseInt(parts[2], 10, 0)
		if err != nil {
			fmt.Println("Invalid number:", parts[2])
			continue
		}

		// Get the operator
		operator := parts[1]

		calculation(int(num1), int(num2), operator, c)

	}

}
