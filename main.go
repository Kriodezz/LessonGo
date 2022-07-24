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

/*
Условия
 if a < b {
      fmt.Println("a меньше, чем b")
 }
If с краткой инструкцией
If может начинаться с инструкции, которая будет выполнена перед проверкой условия.
Переменные, объявленные в этом блоке, доступны только в области видимости,
которая существует до конца if. Пример:
if v := math.Pow(x, n); v < lim {
   // ...
}
else if - через пробел
*/

/*
switch i {
case 0: fmt.Println("Zero")
case 1: fmt.Println("One")
case 2: fmt.Println("Two")
case 3: fmt.Println("Three")
case 4: fmt.Println("Four")
case 5: fmt.Println("Five")
default: fmt.Println("Unknown Number")
}
В пхп по умолчанию продолжается выполнение кода, break - прерывает
в go по умолчанию прерывает, чтобы выполнился след блок
case 42:
	fmt.Println(42)
	fallthrough
Аналог switch (true):
switch {
case 1 <= c && c <= 9:
	fmt.Println("от 1 до 9")
case 100 <= c && c <= 250:
	fmt.Println("от 100 до 250")
}
*/

func main() {
	var x int
	fmt.Scan(&x)
	if (x%400 == 0) || ((x%4 == 0) && (x%100 != 0)) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
