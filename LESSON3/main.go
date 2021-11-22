package main

import (
	"errors"
	"fmt"
	"math"
	"time"
)

func main() {
	err := Calc()
	if err != nil{
		fmt.Println(err)
	}
}

func Calc() error{
	var a, b, res float64
	var op string

	fmt.Print("Введите первое число: ")
	fmt.Scanln(&a)

	fmt.Print("Введите второе число: ")
	fmt.Scanln(&b)

	fmt.Print("Введите арифметическую операцию (+, -, *, /,sqrt, prime): ")
	fmt.Scanln(&op)

	switch op {
	case "+":
		res = a + b
	case "-":
		res = a - b
	case "*":
		res = a * b
	case "/":
		res = a / b
		if b == 0 {
			//fmt.Println("Ошибка на ноль делить нельзя")
			return errors.New("Ошибка на ноль делить нельзя")
		}
	case "sqrt":
		res = math.Sqrt(float64(a))
	case "prime":
		startAt := time.Now() //Время старта
		b := int(math.Floor(a))
		for j := 2; j < b; j++ { //перебераем числа от 2 до <a,так как знаем что 1 не простое число то и перебор делаем с 2
			var flag bool = true
			//ищем корень квадратный от j для сокращение количество перебора
			//если все числа до корня квадратные простые, то и после корня они простые экономим время
			//в коде добавлен простой счетчик выполнения когда проходили полностью j/2  на вываолненеи 100000 понадобилось 13мс
			//после оптимизации стало 6
				m := int(math.Floor(math.Sqrt(float64(j))))
				for i := 2; i <= m; i+=2 {
					if j%i == 0 { //если число делится без остатка на i то оно не простое
						flag = false
						break
					}
				}
			if flag == true {
				fmt.Println(j, "- простое число")
			}
		}
		fmt.Println("Проверка прошла за: ", time.Now().Sub(startAt)) //Высчитываем продолжительность проверки
	default:
		errors.New("Операция выбрана неверно")
		break
	}

	fmt.Printf("Результат выполнения операции: %.4f\n", res)
	return nil
}
