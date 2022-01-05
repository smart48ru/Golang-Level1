package main

import (
	"fmt"
	"math"
	"os"
)
//Value - Структура вводных данных
type Value struct {
	A float64
	B float64
}
//Addition - метод сложения
func (s *Value) Addition() float64 {
	return s.A + s.B
}
//Subtraction - метод вычитания
func (s *Value) Subtraction() float64 {
	return s.A - s.B
}
//Multiplication - метод умножения
func (s *Value) Multiplication() float64 {
	return s.A * s.B
}
//Division - метод деления
func (s *Value) Division() float64 {
	return s.A / s.B
}
//Square - метод кореньквадратный
func (s *Value) Square() float64 {
	return math.Sqrt(float64(s.A))
}
//AreaCalculated - интерфейс для функций калькулятора
type AreaCalculated interface {
	Addition() float64
	Subtraction() float64
	Multiplication() float64
	Division() float64
	Square() float64
}


func main()  {
	var d, e, res float64
	var op string

	fmt.Print("Введите первое число: ")
	fmt.Scanln(&d)

	fmt.Print("Введите второе число: ")
	fmt.Scanln(&e)

	fmt.Print("Введите операцию (+, -, *, /, sqrt, ^, quit): ")
	fmt.Scanln(&op)
	a := &Value{A: d, B: e}
	switch op {
	case "+":
		res = AreaCalculated.Addition(a)
	case "-":
		res = AreaCalculated.Subtraction(a)
	case "*":
		res = AreaCalculated.Multiplication(a)
	case "/":
		if e == 0 {
			fmt.Println("Ошибка на ноль делить нельзя")
			return
		}
		res = AreaCalculated.Division(a)
	case "sqrt":
		res = AreaCalculated.Square(a)
	case "quit":
		os.Exit(1)
	default:
		fmt.Println("Операция выбрана неверно")
		break
	}
	fmt.Printf("Результат выполнения операции: %.2f\n", res)
}