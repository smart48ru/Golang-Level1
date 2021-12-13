package operations

import "fmt"

const OperationTypeMinus OperationType = "-"

type  OperationMinus struct {
	x, y float64
}

func NewOperationMinus() *OperationMinus {
	return &OperationMinus{}
}

func (o *OperationMinus) Calc() (float64, error) {
	return o.x - o.y, nil
}

func (o *OperationMinus) Init(arguments ...float64) error {
	if len(arguments) != 2 {
		return fmt.Errorf("OperationMinus requires 2 arguments")
	}

	o.x = arguments[0]
	o.y = arguments[1]
	return nil
}

func (o *OperationMinus) Match(opType OperationType) bool {
	return opType == OperationTypeMinus
}
