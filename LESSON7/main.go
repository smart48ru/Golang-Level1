package main

import (
	"fmt"
)

func main() {

	/*
		1. С какими интерфейсами мы уже сталкивались в предыдущих уроках?
		Обратите внимание на уроки, в которых мы читали из стандартного ввода и писали в стандартный вывод.

		- fmt.Println
	*/

	var b string
	fmt.Println("Тестовый текст. принимает интерфейс")
	fmt.Println("______________")
	fmt.Println("Введите текст")
	_, err := fmt.Scan(&b)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(b)
}
