package main

import "fmt"

var students = make(map[string]int)

func main() {
	students["Иван"] = 20
	students["Jane"] = 21
	students["Jim"] = 22

	for _, score := range students {
		fmt.Println(score)
	}

	value, ok := students["Иван"] 
	if ok {
		fmt.Println("Иван тут", value)
	} else {
		fmt.Println("Ивана тут нет")
	}
	
	delete(students, "Jane")
	fmt.Println(students)
}