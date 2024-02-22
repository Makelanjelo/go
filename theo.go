package main

import "fmt"

// type User struct {
// 	name     string
// 	age      int64
// 	password string
// }

// func change(u *User) {
// 	u.name = "Kate"
// }

// func (u *User) setName(name1 string) {
// 	u.name = name1
// }

// func (u *User) setAge() bool {
// 	a := u.age
// 	isTrue := false
// 	if a >= 18 {
// 		isTrue = true
// 	} else {
// 		isTrue = false
// 	}
// 	return isTrue
// }

// Интерфейсы

// type Numbers struct {
// 	num1 int
// 	num2 int
// }

// type Interface interface {
// 	Sum() int
// 	Mul() int
// 	Div() float64
// 	Sub() int
// }

// func (n Numbers) Sum() int {
// 	return n.num1 + n.num2
// }
// func (n Numbers) Mul() int {
// 	return n.num1 * n.num2
// }
// func (n Numbers) Div() float64 {
// 	return float64(n.num1) / float64(n.num2)
// }
// func (n Numbers) Sub() int {
// 	return n.num1 - n.num2
// }

func main() {
	fmt.Println("Program run!")
	// user := User{"Jhon", 12, "pass"} // Создане Данных по нашему модулю
	// change(&user)                    // Замена в ячейке памяти
	// user.age = 43                    // Замена без изменения памяти
	// fmt.Println(user)

	// user.setName("Kate")
	// fmt.Println(user)
	// if user.setAge() {
	// 		fmt.Println("Заходи")
	// } else {
	// 		fmt.Println("Стой")
	// }

	// 	var i Interface
	// 	numbers := Numbers{9, 8}
	// 	i = numbers
	// 	fmt.Printf("Сумма чисел %d\n", i.Sum())
	// 	fmt.Printf("Произведение чисел %d\n", i.Mul())
	// 	fmt.Printf("Частное чисел %f\n", i.Div())
	// 	fmt.Printf("Разность чисел %d\n", i.Sub())

}
