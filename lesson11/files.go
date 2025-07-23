package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Создание и запись в файл
	textFile, err := os.Create("test.txt")
	if err != nil {
		fmt.Println("Ошибка открытия файла:", err)
		return
	}
	_, _ = textFile.WriteString("Строка 1\n")
	_, _ = textFile.WriteString("Строка 2\n")
	_, _ = textFile.WriteString("Строка 3\n")
	_, _ = textFile.WriteString("Строка 4\n")
	_, _ = textFile.WriteString("Строка 5\n")
	textFile.Close() // Закрываем после записи

	// Открываем для чтения
	readFile, err := os.Open("test.txt")
	if err != nil {
		fmt.Println("Ошибка открытия файла для чтения:", err)
		return
	}
	defer readFile.Close()

	scanner := bufio.NewScanner(readFile)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Ошибка чтения файла:", err)
	}

	// 2
	if len(os.Args) < 3 {
		fmt.Println("Использование: go run files.go <имя_файла> <количество_строк>")
		return
	}

	fileName := os.Args[1]
	countStr := os.Args[2]

	count, err := strconv.Atoi(countStr)
	if err != nil || count < 1 {
		fmt.Println("Ошибка: количество строк должно быть положительным числом")
		return
	}

	// Создание и запись в файл
	textFile2, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Ошибка создания файла:", err)
		return
	}
	for i := 1; i <= count; i++ {
		_, _ = textFile2.WriteString(fmt.Sprintf("Строка %d\n", i))
	}
	textFile2.Close()

	// Открытие для чтения
	readFile2, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Ошибка открытия файла для чтения:", err)
		return
	}
	defer readFile.Close()

	scanner2 := bufio.NewScanner(readFile2)
	for scanner.Scan() {
		fmt.Println(scanner2.Text())
	}
	if err := scanner2.Err(); err != nil {
		fmt.Println("Ошибка чтения файла:", err)
	}

}