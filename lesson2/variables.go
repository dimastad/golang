package variables
import "fmt"

var name string = "Dimas"
var age int = 35
var weight float64 = 76.9

func main() {
    var isMarried bool = true
    fmt.Println("Hello Go!")
    fmt.Println("Имя: ", name)
    fmt.Println("Возраст: ", age)
    fmt.Println("Вес: ", weight)
    fmt.Println("Женат: ", isMarried)

}