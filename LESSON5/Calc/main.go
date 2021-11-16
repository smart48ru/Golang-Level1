package main

import (
	"fmt"
	"math"
	"os"
	"time"
)

var firstSlice []float64 //слайс для первых чисел
var secondSlice []float64 //слайс для вторых чисел
var operatorSlice []string //создаем слайс операторов
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
	var add bool //фиксировать или нет результат в салйс сделал для того чтобы при вводе prime не выводилось в этот слайс
	fmt.Print("Введите первое число: ")
	fmt.Scanln(&a)

	fmt.Print("Введите второе число: ")
	fmt.Scanln(&b)

	fmt.Print("Введите операцию (+, -, *, /,sqrt, ^, prime , result, quit): ")
	fmt.Scanln(&op)

	add = true
	switch op {
	case "+":
		res = a + b
	case "-":
		res = a - b
	case "*":
		res = a * b
	case "/":
		if b == 0 {
			fmt.Println("Ошибка на ноль делить нельзя")
		}
		res = a / b
	case "sqrt":
		res = math.Sqrt(float64(a))
	case "^":
		res = math.Pow(float64(a),float64(b))
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
				prePrimeSlice = append(prePrimeSlice, float64(j))//добавляем результат в слайс простых чисел
			}
		}
		primeSlice = append(primeSlice, prePrimeSlice)
		fmt.Println("Проверка прошла за: ", time.Now().Sub(startAt)) //Высчитываем продолжительность проверки
	case "quit":
		os.Exit(1)
	case "result":
		for id, opers := range operationsSlice {

			a = firstSlice[id]
			b = secondSlice[id]
			op = operatorSlice[id]
			fmt.Println("Результат операции", id,a,op, b, "=",opers)
		}

		for id, val := range primeSlice {
			fmt.Println("Результат операции с простми чисел", id,"=",val)
		}
		add = false
	default:
		fmt.Println("Операция выбрана неверно")
		break
	}
	if add == true {
		firstSlice = append(firstSlice, a)             //добавляем перове число в слайс
		secondSlice = append(secondSlice, b)           //добавляем перове число в слайс
		operatorSlice = append(operatorSlice, op)      //добавляем оперант в слайс
		operationsSlice = append(operationsSlice, res) //добавляем результат в слайс операций
	}
	fmt.Printf("Результат выполнения операции: %.4f\n", res)
}