package main

import (
	"fmt"
	"math"
	"os"
)

//Value - Структура вводных данных
type Value struct {
	a float64
	b float64
}

//Addition - метод сложения
func (s *Value) Addition() float64 {
	return s.a + s.b
}

//Subtraction - метод вычитания
func (s *Value) Subtraction() float64 {
	return s.a - s.b
}

//Multiplication - метод умножения
func (s *Value) Multiplication() float64 {
	return s.a * s.b
}

//Division - метод деления
func (s *Value) Division() float64 {
	return s.a / s.b
}

//Square - метод кореньквадратный
func (s *Value) Square() float64 {
	return math.Sqrt(float64(s.a))
}

//AreaCalculated - интерфейс для функций калькулятора
type AreaCalculated interface {
	Addition() float64
	Subtraction() float64
	Multiplication() float64
	Division() float64
	Square() float64
}

func main() {
	var d, e, res float64
	var op string

	fmt.Print("Введите первое число: ")
	fmt.Scanln(&d)

	fmt.Print("Введите второе число: ")
	fmt.Scanln(&e)

	fmt.Print("Введите операцию (+, -, *, /, sqrt, ^, quit): ")
	fmt.Scanln(&op)
	val := &Value{a: d, b: e}
	switch op {
	case "+":
		res = AreaCalculated.Addition(val)
	case "-":
		res = AreaCalculated.Subtraction(val)
	case "*":
		res = AreaCalculated.Multiplication(val)
	case "/":
		if e == 0 {
			fmt.Println("Ошибка на ноль делить нельзя")
			return
		}
		res = AreaCalculated.Division(val)
	case "sqrt":
		res = AreaCalculated.Square(val)
	case "quit":
		os.Exit(1)
	default:
		fmt.Println("Операция выбрана неверно")
		os.Exit(1)
	}
	fmt.Printf("Результат выполнения операции: %.2f\n", res)
}
