package main

import (
	"fmt"
	"math"
	"os"
	"time"
)

var operationsSlice []float64 //создаем слайс для сохранение ответов операций
var primeSlice [][]float64 // создаем слайс для сохранения ответов простых чисел



func main() {
	for {
		Calc() //Вынес в отдельную функцию для простоты перезапуска
	}
}

func Calc() {
	var a, b, res float64
	var op string

	fmt.Print("Введите первое число: ")
	fmt.Scanln(&a)

	fmt.Print("Введите второе число: ")
	fmt.Scanln(&b)

	fmt.Print("Введите операцию (+, -, *, /,sqrt, ^, prime , result, quit): ")
	fmt.Scanln(&op)


	switch op {
	case "+":
		res = a + b

	case "-":
		res = a - b
		//operationsSlice = append(operationsSlice, res)//добавляем результат в слайс операций
	case "*":
		res = a * b
		//operationsSlice = append(operationsSlice, res)//добавляем результат в слайс операций
	case "/":
		if b == 0 {
			fmt.Println("Ошибка на ноль делить нельзя")
		}
		res = a / b
		//operationsSlice = append(operationsSlice, res)//добавляем результат в слайс операций
	case "sqrt":
		res = math.Sqrt(float64(a))
		//operationsSlice = append(operationsSlice, res)//добавляем результат в слайс операций
	case "^":
		res = math.Pow(float64(a),float64(b))
		//operationsSlice = append(operationsSlice, res)//добавляем результат в слайс операций
	case "prime":
		startAt := time.Now() //Время старта
		var prePrimeSlice []float64
		b := int(math.Floor(a))
		for j := 2; j < b; j++ { //перебераем числа от 2 до <a,так как знаем что 1 не простое число то и перебор делаем с 2
			var flag bool = true
			//ищем корень квадратный от j для сокращение количество перебора
			//если все числа до корня квадратные простые, то и после корня они простые экономим время
			//в коде добавлен простой счетчик выполнения когда проходили полностью j/2  на вываолненеи 100000 понадобилось 13мс
			//после оптимизации стало 6
			if j%2 == 0 { // Оптимизация дала большое ускорение 87.101541ms
				flag = false
			} else {
				m := int(math.Floor(math.Sqrt(float64(j))))
				for i := 2; i <= m; i++ {
					if j%i == 0 { //если число делится без остатка на i то оно не простое
						flag = false
						break
					}
				}
			}
			if flag == true {

				//fmt.Println(j, "- простое число")
				prePrimeSlice = append(prePrimeSlice, float64(j))//добавляем результат в слайс простых чисел
			}
		}
		primeSlice = append(primeSlice, prePrimeSlice)//ДА добовление в массив можно вынести за case. неподумал об этом
		fmt.Println("Проверка прошла за: ", time.Now().Sub(startAt)) //Высчитываем продолжительность проверки
	case "quit":
			os.Exit(1)
	case "result":
		for id, val := range operationsSlice {
			fmt.Println("Результат операции", id,"=",val)
		}

		for id, val := range primeSlice {
			fmt.Println("Результат операции с простми чисел", id,"=",val)
		}
	default:
		fmt.Println("Операция выбрана неверно")
		break
	}
	operationsSlice = append(operationsSlice, res)//добавляем результат в слайс операций
	fmt.Printf("Результат выполнения операции: %.4f\n", res)
}
