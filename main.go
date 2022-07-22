package main

import "fmt"

//func main() {
//	fmt.Println(string("I like Go!"[0])) //Обращение к символу
//}

/*
 Обьявление переменных
 var hello string
 Обьявление нескольких переменных
 var oneVar, twoVar, threeVar string
 Типы можно не указывать, например var hello = "Hello" или var x = 10
 Когда объявляется переменная, она автоматически содержит значение по умолчанию для своего типа:
 0 для int, 0.0 для float, false для bool, пустая строка для string, nil для указателя и т.д.
 Краткое объявление переменных (доступно внутри функций): a := 5 (аналог var a = 5)
 Также можно объявить сразу несколько переменных в одном блоке var:
 var (
	 name string = "Dima"
	 age int = 23
 )
*/

//func main() {
//	var hello string
//	hello = "Hello Go!"
//	var x = 2022
//	fmt.Println(hello)
//	fmt.Println(x)
//}

/*
 Для хранения символов можно использовать int32/rune. Здесь используются одинарные кавычки.
 Компилятор определяет код буквы в unicode и присваивает его переменной symbol.
 То есть мы не храним никакую 'c', а храним лишь число 99.
 Функция string() из переданного в него числа 99 делает строку 'c'.
 var symbol int32 = 'c'
 fmt.Println(string(symbol))
*/

/*
 Ввод данных. присвоение данных переменным, вывод данных
*/
//func main() {
//	var name string
//	var age int
//	fmt.Print("Введите имя: ")
//	fmt.Scan(&name)
//	fmt.Print("Введите возраст: ")
//	fmt.Scan(&age)
//
//	fmt.Println(name, age)
//}

// аналог echo "My name is $name and I am $age years old.";
//func main() {
//	name := "Ivan"
//	age := 27
//	fmt.Println("My name is", name, "and I am", age, "years old.")
//}

/*
 Константы
 обьявление const pi float64 = 3.1415
 обьявление нескольких констант
 const (
	a int = 45
	b float32 = 3.3
 )
*/

/*
Идентификатор iota начинает счетчик с 0 для блока констант
const (
	Sunday = iota //0
	Monday
	Tuesday
	Wednesday  //3
	Thursday
	Friday
	Saturday  //6
)
*/

func main() {
	fmt.Println("Hello")
}
