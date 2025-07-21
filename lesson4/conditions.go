package main

import "fmt"

func main() {
	var number int = 0

	if number%2 == 0 {
		fmt.Println("Число четное")
	} else {
		fmt.Println("Число нечетное")
	}

	day := "Monday"

	switch day {
	case "Monday":
		fmt.Println("Понедельник")
	case "Tuesday":
		fmt.Println("Вторник")
	case "Wednesday":
		fmt.Println("Среда")
	case "Thursday":
		fmt.Println("Четверг")
	case "Friday":
		fmt.Println("Пятница")
	case "Saturday":
		fmt.Println("Суббота")
	case "Sunday":
		fmt.Println("Воскресенье")
	}
}
