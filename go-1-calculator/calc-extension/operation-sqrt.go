package calc_extension

import (
	"fmt"
	"gb/go-1-11/calculator/operations"
	"math"
)

const OperationTypeSqrt operations.OperationType = "sqrt"

type OperationSqrt struct {
	x float64
}

func NewOperationSqrt() *OperationSqrt {
	return &OperationSqrt{}
}

func (o *OperationSqrt) Calc() (float64, error) {
	return math.Sqrt(o.x), nil
}

func (o *OperationSqrt) Init(arguments ...float64) error {
	if len(arguments) < 1 {
		return fmt.Errorf("OperationSqrt requires 1 argument")
	}

	o.x = arguments[0]
	return nil
}

func (o *OperationSqrt) Match(opType operations.OperationType) bool {
	return opType == OperationTypeSqrt
}
