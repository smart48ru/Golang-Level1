package main

import (
	"fmt"
	"time"
)
//Структура глобальной map для кеша. Объявляю глобально и не очищаю его для сохранения результатов предыдущей генерации

type cache struct {
	Map map[int]int
}


//Создаем метод мап через структуру cache. Незнаю правильно ли так делать.
var (
	m = cache{Map: map[int]int{}}

)

func main() {
	for {
		var n int
		fmt.Print("Введите число: ")
		fmt.Scanln(&n)
		startAt := time.Now() //Время старта
		fmt.Println("------------ Время выполнения только для того что бы понять разницу ------------")

		fmt.Println(fibMapCache(n, m), "= результат от числа",n )
		fmt.Println("Проверка с кешем прошла за: ", time.Now().Sub(startAt))

		startAt = time.Now()
		fmt.Println(fibbonachi(n),"= результат от числа",n )
		fmt.Println("Проверка проверка рекурсия прошла за: ", time.Now().Sub(startAt))
	}
}
/*
fibMapCache - Функция Фиббоначи с кешем мап.
если x < 2 то добавляем в map текущие значения x.
Заполняем мапу через рекурсию. Если x-1 не существует то заполняем его рекурсией до точки выхода.
возвращаем значение по формуле из мапы m.Map[x-1] + m.Map[x-2]
*/
func fibMapCache(x int, m cache) int {
	if x < 2 {
		m.Map[x] = x
		return x
	}
	_, ok := m.Map[x-1]
	if !ok {
		m.Map[x-1] = fibMapCache(x-1, m)
	}
	return m.Map[x-1] + m.Map[x-2]
}

/*
fibbonachi - Функция Фиббоначи чистая рекурсия
*/
func fibbonachi(n int) int {
	if n <2 {
		return n
	}
	return fibbonachi(n-1) + fibbonachi(n-2)
}
