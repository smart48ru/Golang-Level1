package main

<<<<<<< HEAD
import "fmt"

func main() {
	/*
		a. В каких случаях необходима была явная передача указателя в качестве входных
		параметров и возвращаемых результатов или в качестве приёмника в методах?

		- в случае елси нам не нужно создавать(клонировать) новый объект в памяти.
	*/

	answer := 42
	fmt.Println(&answer)
	//Передаем по ссылки получаем тот же адрес памяти
	address := &answer
	fmt.Println(address)
	//Создаем новые переменные получаем 2 адреса памяти
	newaddress := answer
	fmt.Println(&newaddress)

	/*
	   b. В каких случаях мы фактически имеем дело с указателями при передаче параметров, хотя явно их не указываем

	   - когда объявляем тип переменной (допустим int)
	*/

	/*
		2. Для арифметического умножения и разыменования указателей в Go используется один и тот же символ —
		оператор (*). Как вы думаете, как компилятор Go понимает, в каких случаях в выражении имеется в виду умножение, а в каких — разыменование указателя?

		- при арефмитическом умнажении у меня не получилось сделать ссылку. invalid indirect of v (type int
			в умножение есть множимое и множитель между ними есть оператор (*)
	*/
	fmt.Println("2-----2----2")
	v := 5
	d := 2

	res := v*d
	fmt.Println(res)
=======
import (
	"fmt"
	"math"
)



type Square struct {
	Side float32
}

func (s *Square) CalculateArea() float32 {
	return s.Side * s.Side
}

type Circle struct {
	Radius float32
}

func (c *Circle) CalculateArea() float32 {
	return math.Pi * c.Radius * c.Radius
}

type AreaCalculator interface {
	CalculateArea() float32
}

func main() {
	square := &Square{Side: 4}
	circle := &Circle{Radius: 4}

	for _, shape := range []AreaCalculator{square, circle} {
		fmt.Println(shape.CalculateArea())
	}
>>>>>>> main
}