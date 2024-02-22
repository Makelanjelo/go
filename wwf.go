package main

import (
	"fmt"
	// "os"
)

func main() {
	fmt.Println("Programm run!")

	// Работа с файлами

	// data := []byte("Text to file")

	// e := os.WriteFile("text.txt", data, 0600)

	// if e != nil {
	// 	fmt.Println("Не могу создать файл\n", e)
	// }

	// file_data, err := os.ReadFile("text.txt") // Открытие файла

	// if err != nil { // Проверка на ошибку
	// 	fmt.Println("Не могу прочитать фаил\n", err)
	// }

	// fmt.Println(string(file_data))

	// os.Remove("text.txt")

	// file, err := os.Create("text.txt")

	// if err != nil {
	// 	println("Ошибка", err)
	// 	os.Exit(1)
	// }
	// defer file.Close()

	// data := "Text to file"

	// file.WriteString(data)

	// file_data, _ := os.ReadFile(file.Name())

	// fmt.Println(string(file_data))

	i := 0
LOOP:
	if i > 50 {
		goto END
	}

	fmt.Println(i)
	i++
	goto LOOP

END:
}
