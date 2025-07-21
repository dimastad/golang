package main
import "fmt"

var price1 int = 100
var price2 int = 200


func main() {
	sum := price1 + price2
	sub := price1 - price2
	mult := price1 * price2
	div := price1 / price2
	mod := price1 % price2

	fmt.Println("Сумма: ", sum)
	fmt.Println("Разность: ", sub)
	fmt.Println("Произведение: ", mult)
	fmt.Println("Частное: ", div)
	fmt.Println("Остаток: ", mod)
}