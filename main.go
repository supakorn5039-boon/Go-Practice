package main

import "fmt"

var numInt, numInt2 int = 1000, 2000

var msg string = "Hello"

func main() {
	numberfloat := 25.4
	fmt.Println(numInt)
	fmt.Println(numInt2)
	fmt.Println(numberfloat)
	fmt.Println(msg)

	fmt.Println(numInt + numInt2)
	fmt.Println(float64(numInt) + numberfloat)
	fmt.Println(msg + "world")
	fmt.Println("my money =", numInt)
}
