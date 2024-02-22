package main //Пакет в котором находится фаил
import (
	"fmt"
	// "sort"
	// "math"
)

// Функции
// func first() { // f1
// 	fmt.Println("Hello from first function")
// }

// func sum(a int, b int) int { // Возврат функции с элементом f2
// 	result := a + b
// 	return result
// }

// func math_func(a int, b int) (sum int, sub int, mul int, div int) { // Возврат функции с несколькими элементами f3
// 	sum = a + b
// 	sub = a - b
// 	mul = a * b
// 	div = a / b
// 	return
// }

// func test(someFunction func(int) int) { // Обращение к функции внутри функции f5
// 	fmt.Println(someFunction(25))
// }

// func test(x string) func() { // Возврат функции (Обязательно без имени) f6
// 	return func() {
// 		fmt.Println(x)
// 	}
// }

// Точка фхода в программу

func main() { //Ты охренеешь, но это функция

	fmt.Println("Start programm!")
	// fmt.Println(math.Sqrt(9)) // Квадратный корень
	// var age int8 = 23 // Целочисленное, тоже самое age := 32
	// var mounth uint8 = 5 // Положительное целочисленное := работает
	// var number float64 = 275.672 // С п.т := работает
	// var name string = "Денис" // Строка := работает
	// var name string
	// var age uint8
	// fmt.Println("What is your name?")
	// fmt.Scan(&name) // Прием данных от пользователя
	// fmt.Println("Hello " + name + "!")
	// fmt.Println("How old are you?")
	// fmt.Scan(&age)
	// fmt.Println("You are " + fmt.Sprint(age) + " years!") // Изменение типа данных, но можно через ","
	// fmt.Println(len(name)) // Длинна строки
	// Массив arr := []int{1, 2, 3, 4, 5}

	// Структура if/else

	// num := 0
	// if num > 0 {
	// 	fmt.Println("Number > 0")
	// } else if num < 0 {
	// 	fmt.Println("Number < 0")
	// } else {
	// 	fmt.Println("Number = 0")
	// }

	// Циклы

	// for i := 0; i < 5; i++ { // Переменна; Условие пока; Действие. Break and continue like python
	// 	fmt.Printf("Hello %d\n", i)
	// }

	// arr := []int{1, 2, 3, 4, 5}

	// for index, element := range arr { // Range loop
	// 	fmt.Printf("Index: %d Element: %d\n", index, element)
	// }
	// matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}} // Двойной массив

	// for _, row := range matrix {
	// 	for _, col := range row {
	// 		fmt.Printf("%d ", col)
	// 	}

	// 	fmt.Println()
	// }

	// Конструкции switch

	// name := "Jhon"

	// switch name {

	// case "Jordan":
	// 	fmt.Println("Hello Jordan!")
	//  fallthrough // Если хотим игнорировать встроенный break

	// case "Kate":
	// 	fmt.Println("Hello Kate!")

	// case "Jhon":
	// 	fmt.Println("Hello Jhon!")

	// default: // Дефолтное значение
	// 	fmt.Println("I don't know you!")
	// }

	// Логические операторы и математические выражения такие же как и в python: && - and,  || - or, != - not.

	// Массивы

	// var names [3]string // Так не надо
	// names[0] = "Jordan"
	// names[2] = "Kate"
	// names[1] = "Emma"

	// names := [3]string{"Jordan", "Kate", "Emma"} // Так надо
	// fmt.Println(names)

	// Многомерный массив

	// array := [3][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}
	// fmt.Println(array)
	// fmt.Println(array[1][2])

	// Срезы

	// slice := []int{3, 1, 2, 5, 7, 4, 8, 3}
	// slice = append(slice, 0) // Добавление в конец
	// slice[0] = 11 // Замена элемента
	// sort.Ints(slice)
	// fmt.Println(slice)

	// for i := 0; i < len(slice); i++ { // Лучше через range
	// 	fmt.Println(slice[i])
	// }

	// for _, el := range slice {
	// 	fmt.Printf("%d\n", el)
	// }

	// Форматирование строк: fmt.Printf("randomString %d" , randomVar), %d for int, %s or string, %f for float, %t for boolean

	// Карты - структуры где значение хранится под ключем, на словари похоже. Данные сортируются автоматически. Не использовать если порядок критичен.

	// var money map[string]int = map[string]int{ // лучше делать так // money := map[string]int {
	// 	"dollars": 1000,
	// 	"euros":   2000,
	// 	"apples":  3,
	// }

	// money["dollars"] = 5000 // Изменение
	// delete(money, "aplles") // Удаление
	// el, status := money["dollars"] // Проверка на наличие элемента
	// fmt.Println((money)) // Для вывода элемента например: (money["dollars"])

	// Вызов функции (смотри функции)

	//first() // f1

	// sum(5, 6) // f2

	// s, su, m, d := math_func(3, 5) // f3
	// fmt.Println(s, su, m, d)

	// Создание вложенной функцции

	// a := func() { // f4
	// 	fmt.Println("Hello!")
	// }
	// a()

	// square := func(x int) int { // f5
	// 	return x * x
	// }

	// test(square)

	// test("Hello!")() // a := test("Hello!") // f6 Два метода возврата функции в функции \n // a()

	// Указатели - код элемента в памяти (* - оператор разименование, для обращение к элементу по коду ячейки)

	// a := 5
	// pointer := &a
	// fmt.Println((*pointer))
}
