package main

import "fmt"

func sayHello(name string) {
	fmt.Println("Hello,", name)
}

func sum(a int, b int) int {
	return a + b
}

func main() {
	sayHello("Dima")
	
	fmt.Println("Sum of 3 and 5:", sum(3, 5))
}