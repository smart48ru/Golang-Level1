package operations

import "fmt"

const OperationTypePlus OperationType = "+"

type OperationPlus struct {
	x, y float64
}

func NewOperationPlus() *OperationPlus {
	fmt.Println("NewOperationPlus")
	return &OperationPlus{}
}

func (o *OperationPlus) Calc() (float64, error) {
	fmt.Println("Calc")
	return o.x + o.y, nil
}

func (o *OperationPlus) Match(opType OperationType) bool {
	fmt.Println("Match")
	return opType == OperationTypePlus
}

func (o *OperationPlus) Init(arguments ...float64) error {
	fmt.Println("Init")
	if len(arguments) != 2 {
		return fmt.Errorf("OperationPlus requires 2 arguments")
	}
	o.x = arguments[0]
	o.y = arguments[1]
	return nil
}
