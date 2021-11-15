package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	slice := sliceReaderWriter()
	fmt.Println("           Начальный массив : ", slice)
	fmt.Println("     Отсортированный массив : ", InsertionSort(slice))
}

/*
valCheck
Функция проверки число или нет.
Получает стринг возвращает буливое значение
false - не число
true - число
*/
func valCheck(val string) bool {
	if _, err := strconv.Atoi(val); err == nil {
		return true
	} else {
		return false
	}
}

/*
sliceReaderWriter
Функция читает ввод с консоли собирает все в слайс и если введено не число то просит переввети
*/
func sliceReaderWriter() (slice []int) {
	var arr []int
	scaner := bufio.NewReader(os.Stdin)
	fmt.Print("Введите числа массива через пробел: ")
	str, _ := scaner.ReadString('\n')
	txt := strings.Fields(str)
	for id, val := range txt {
		if valCheck(val) == true {
			intval, _ := strconv.Atoi(val)
			arr = append(arr, intval)
		} else {
			var a string
			fmt.Println("В ячейку № -", id, "введено неверное значение -", val)
			fmt.Print("Введите число , (d) для удаления или любую букву для начала ввода заново: \n")
			fmt.Scanln(&a)
			if valCheck(a) == true {
				intval, _ := strconv.Atoi(a)
				arr = append(arr, intval)
			} else if a == "d" {
				fmt.Print("ячейка № - ", id, " УДАЛЕНА")
			} else {
				fmt.Println("!!Вы опять ввели не число, программа начнется заново!!")
				sliceReaderWriter()
				return
			}
		}
	}
	return arr
}

/*
InsertionSort
Функция сортирует переданный ей слайс и возвращает отсортерованный
Комментарии в функциях пока пишу больше для себя, что бы потом незабыть
*/
func InsertionSort(arr []int) (arrsort []int) {
	arrsort = make([]int, len(arr))     //создаем слайс длинно как у передаваемого функции слайса
	copy(arrsort, arr)                  // сопируем передаваемый слайс в новый слайс
	for i := 1; i < len(arrsort); i++ { //Итрируемся по слайсу начиная с 1 элемента
		x := arrsort[i]                         //присваемаем переменной x значение текущей итеррируемой ячейки
		j := i                                  //присваиваем j номер итерации
		for ; j >= 1 && arrsort[j-1] > x; j-- { //начинаем от j в обратную сторону
			//пока она не станет большще или равно 1 и значение текущая ячейка минус 1 не станет больше x которое мы сохранили в начеле второй итеррации
			arrsort[j] = arrsort[j-1] //присваиваем текущей итерируемой ячейки значение предыдущей
		}
		arrsort[j] = x //записываем в текущую ячейку значение x которое мы сохранили в начеле второй итеррации
	}
	return arrsort
}
