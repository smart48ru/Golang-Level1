package calculator

import (
	"fmt"
	"gb/go-1-11/calculator/operations"
	"strconv"
)

func readInput() (a, b float64, op operations.OperationType, err error) {
	var aStr, bStr string
	fmt.Println("Введите {операнд-1} {действие} {операнд-2} ")

	if _, err = fmt.Scanf("%s %s %s", &aStr, &op, &bStr); err != nil {
		err = fmt.Errorf("не верный формат ввода")
		return
	}

	if a, err = strconv.ParseFloat(aStr, 64); err != nil {
		err = fmt.Errorf("первый операнд не является числом")
		return
	}

	if b, err = strconv.ParseFloat(bStr, 64); err != nil {
		err = fmt.Errorf("второй операнд не является числом")
		return
	}

	return
}
