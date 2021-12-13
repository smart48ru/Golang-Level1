package calculator

import (
	"fmt"
	"gb/go-1-11/calculator/operations"
)

type Calculator struct {
	operations []operations.Operation
}

func NewBaseCalculator() *Calculator {
	fmt.Println("NewBaseCalculator")
	return &Calculator{operations: []operations.Operation{
		operations.NewOperationPlus(),
		operations.NewOperationMinus(),
		operations.NewOperationMultiply(),
		operations.NewOperationDivide(),
		operations.NewOperationExit(),
	}}
}

func (c *Calculator) AppendOperation(ops ...operations.Operation) *Calculator {
	c.operations = append(c.operations, ops...)
	return c
}

func (c *Calculator) Calc() (float64, error) {
	x, y, inputOperation, err := readInput()
	if err != nil {
		return 0, err
	}
	for _, op := range c.operations {
		if op.Match(inputOperation) {
			if err := op.Init(x, y); err != nil {
				return 0, fmt.Errorf("operation %s init failed: %w", inputOperation, err)
			}
			return op.Calc()
		}
	}
	return 0, fmt.Errorf("unsuported operation")
}
