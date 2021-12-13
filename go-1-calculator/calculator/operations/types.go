package operations

type OperationType string

type Operation interface {
	Calc() (float64, error)
	Init(arguments ...float64) error
	Match(operationType OperationType) bool
}
