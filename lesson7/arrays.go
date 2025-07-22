package main
import "fmt"

func array() [5]int {
	var array [5]int
	for i := 0; i < 5; i++ {
		array[i] = i + 1
	}
	return array
}

var sliceArr []int

var arr2 = [6]int{2, 4, 6, 8, 10, 12}
var slice2 = arr2[1:4]

func main() {
	
	numbers := array()
	fmt.Println(numbers)

	
	sliceArr = append(sliceArr, 10, 20, 30)
	fmt.Println(len(sliceArr), cap(sliceArr))
	fmt.Println(slice2)
}
