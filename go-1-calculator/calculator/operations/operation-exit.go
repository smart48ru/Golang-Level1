package operations

import (
	"errors"
)

var ErrExit = errors.New("exit")

const OperationTypeExit OperationType = "exit"

type OperationExit struct{}

func NewOperationExit() *OperationExit {
	return &OperationExit{}
}

func (o *OperationExit) Calc() (float64, error) {
	return 0, ErrExit
}

func (o *OperationExit) Init(arguments ...float64) error {
	return nil
}

func (o *OperationExit) Match(opType OperationType) bool {
	return opType == OperationTypeExit
}
