package main

import (
	"errors"
	"fmt"
	calc_extension "gb/go-1-11/calc-extension"
	"gb/go-1-11/calculator"
	"gb/go-1-11/calculator/operations"
	"strings"
)

const (
	ActionTry  = "try"
	ActionExit = "exit"
)

func main() {
	calc := calculator.NewBaseCalculator()
	calc.AppendOperation(calc_extension.NewOperationSqrt())

	for {
		result, err := calc.Calc()
		if err != nil {
			if errors.Is(err, operations.ErrExit) {
				fmt.Println("bye")
				return
			}

			fmt.Println("calc err: ", err.Error())

			var action string
			fmt.Printf("to try again enter `%s` or `%s` for exit\n", ActionTry, ActionExit)
			if _, err := fmt.Scan(&action); err != nil {
				fmt.Printf("could not continue because of scan err: %s\n", err)
				return
			}
			// Sanitizing
			action = strings.TrimSpace(action)
			action = strings.ToLower(action)
			if action == ActionTry {
				continue
			} else if action == ActionExit {
				fmt.Println("bye")
				return
			}

			return
		}

		fmt.Println("Result = ", result)
	}
}
