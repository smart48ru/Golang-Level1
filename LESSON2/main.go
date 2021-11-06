package main

import (
	"fmt"
	"math"
)

func main() {
	var task string
	fmt.Println(`Урок №2! Введите номер домашего задания 1-3 или "q" для выхода`)
	_, err := fmt.Scanln(&task)
	if err != nil {
		fmt.Println("Ошибка !!", err.Error())
	}
	switch task {
	case "1":
		task1() //Запуск функции задание 1
	case "2":
		task2() //Запуск функции задание 2
	case "3":
		task3() //Запуск функции задание 3
	case "q":
		return
	default:
		fmt.Println("Вы ввели не верный номер задания")
		main() //введен неверный номер задания запускаем заново
	}

}

func task1() {
	var err error
	fmt.Println("Задание 1 - Программа для вычисления площади прямоугольника.")
	fmt.Println("Для вычисления площади прямоугольника введите длинну стороны a и b")
	var a, b float32 //так как мы не знаем будет введено целое число или дробное объявляем float32
	fmt.Println("Введите длинну стороны прямоугольника a: ")
	_, err = fmt.Scanln(&a)
	if err != nil {
		fmt.Println("Ошибка !!", err.Error())
	}
	fmt.Println("Введите длинну стороны прямоугольника b: ")
	_, err = fmt.Scanln(&b)
	if err != nil {
		fmt.Println("Ошибка !!", err.Error())
	}
	fmt.Printf("Площадь прямоугольника: %f\n", a*b)
}

func task2() {
	fmt.Println("Задание 2 - Программа для вычисления диаметра и длины окружности по заданной площади круга")
	fmt.Println("Введите площадь круга")
	var s float64 //для расчетов корня квадратного нам нужен тип float64
	_, err := fmt.Scanln(&s)
	if err != nil {
		fmt.Println("Ошибка !!", err.Error())
	}
	result_d := math.Sqrt(s/math.Pi) * 2 //Если я не ошибаюсь в формуле D=√S/π
	result_p := result_d * math.Pi       // Зная радиус высчитываем по формуле  P=D*π
	fmt.Printf("Диаметр окружности: %f\n", result_d)
	fmt.Printf("Длинна окружности: %f\n", result_p)
}

func task3() {
	fmt.Println("Задание 3 - Выведите цифры, соответствующие количество сотен, десятков и единиц  в этом числе.")
	fmt.Println("Введите трехзначное число")
	var numbers string
	_, err := fmt.Scanln(&numbers)
	if err != nil {
		fmt.Println("Ошибка !!", err.Error())
	}
	if len(numbers) == 3 { //проверяем длинну числа должна равняться 3
		fmt.Println("Число:", numbers)
		fmt.Println("Сотни: ", string(numbers[0]))
		fmt.Println("Десятки: ", string(numbers[1]))
		fmt.Println("Еденицы: ", string(numbers[2]))
	} else {
		fmt.Println("Вы ввели не верное згначение. !! Число не трехзначное !!")
	}
}
