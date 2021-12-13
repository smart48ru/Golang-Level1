package operations

import "fmt"

const OperationTypeMultiply OperationType = "*"

type OperationMultiply struct {
	x, y float64
}

func NewOperationMultiply() *OperationMultiply {
	return &OperationMultiply{}
}

func (o *OperationMultiply) Calc() (float64, error) {
	return o.x * o.y, nil
}

func (o *OperationMultiply) Match(opType OperationType) bool {
	return opType == OperationTypeMultiply
}

func (o *OperationMultiply) Init(arguments ...float64) error {
	if len(arguments) != 2 {
		return fmt.Errorf("OperationMultiply requires 2 arguments")
	}

	o.x = arguments[0]
	o.y = arguments[1]
	return nil
}
