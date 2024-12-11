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

func leaseExtend(T int, c calculatorproxy.CalculatorProxy) {
	_, status := c.LeaseExtend(T)
	if status == "ok" {
		// Display the result
		fmt.Println("Command to extend lease by", T, "seconds!")
		return
	} else {
		fmt.Println("operation not available. status: ", status)
		return
	}
}

func leaseTypeSet(lease string, c calculatorproxy.CalculatorProxy) {
	_, status := c.LeaseTypeSet(lease)
	if status == "ok" {
		// Display the result
		fmt.Println("Command sent correctly")
		return
	} else {
		fmt.Println("operation not available. status: ", status)
		return
	}

}

func calculation(num1 int, num2 int, operator string, c calculatorproxy.CalculatorProxy) {
	var result int
	var status string
	switch operator {
	case "+":
		result, status = c.Sum(num1, num2)
	case "-":
		result, status = c.Sub(num1, num2)
	case "*":
		result, status = c.Mul(num1, num2)
	case "/":
		result, status = c.Div(num1, num2)
	default:
		fmt.Println("Invalid operator. Use one of: +, -, *, /")
		return
	}

	if status == "ok" {
		// Display the result
		fmt.Println("Result:", result)
		return
	} else {
		fmt.Println("operation not available. status: ", status)
		return
	}

}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Welcome to the Calculator!")
	fmt.Println("Enter your calculation in the format: number1 operator number2 (e.g., 12 + 5)")
	fmt.Println("Type 'exit' to quit, 'extend_lease' to keep using calculator, 'lease_type_[x]' to set the type of leasing, [x] can be 0, 1 or 2 ")

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

		if input == "lease_type_0" {
			// nesse tipo de invocacao, o lease eh renovado a cada chamada do obj remoto.
			fmt.Println("TIPO 0")
			leaseTypeSet("lease_type_0", c)

		} else if input == "lease_type_1" {
			// nesse tipo de invocacao, o lease somente eh renovado por uma chamada especifica do cliente: leaseExtend()
			fmt.Println("TIPO 1")
			leaseTypeSet("lease_type_1", c)

		} else if input == "lease_type_2" {
			fmt.Println("TIPO 2")
			leaseTypeSet("lease_type_2", c)

			/**
			The distributed object middleware informs the client of a lease’s
			upcoming expiration, allowing the client to specify an extension
			period. The client does not have to keep track of lease expiration
			itself, and thus the logic to manage the lifecycle of its remote
			objects becomes simpler.

			On the other hand, as network communication is unreliable, lease expiration messages might get lost and
			remote objects might be destroyed unintentionally. A further
			liability is that clients need to be able to handle such messages, which typically requires them to provide callback remote objects,
			so they have to be servers, too.
			*/

		} else {
			fmt.Println("AAAAAAAAAA")
			// Split the input
			parts := strings.Split(input, " ")

			if (len(parts) != 3) && (len(parts) != 2) {
				fmt.Println("Invalid input format. Use: number1 operator number2")
				continue
			}

			if len(parts) == 2 {
				T, err := strconv.ParseInt(parts[1], 10, 0)
				if err != nil {
					fmt.Println("Invalid number:", parts[1])
					continue
				}
				if parts[0] == "extend_lease" {
					leaseExtend(int(T), c)
				}
			} else {
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
	}

}
