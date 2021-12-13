package operations

import "fmt"

const OperationTypeDivide OperationType = "/"

type OperationDivide struct {
	x, y float64
}

func NewOperationDivide() *OperationDivide {
	return &OperationDivide{}
}

func (o *OperationDivide) Calc() (float64, error) {
	if o.y == 0 {
		return 0, fmt.Errorf("could not divide by zero")
	}
	return o.x / o.y, nil
}

func (o *OperationDivide) Match(opType OperationType) bool {
	return opType == OperationTypeDivide
}

func (o *OperationDivide) Init(arguments ...float64) error {
	if len(arguments) != 2 {
		return fmt.Errorf("OperationDivide requires 2 arguments")
	}

	o.x = arguments[0]
	o.y = arguments[1]
	return nil
}
