package main

import "fmt"

func Hello() {
	fmt.Println("Hello Worlds!")
}

// value then types
func Plus1(a int, b int) {
	result := a + b
	fmt.Println(result)
}

func Plus2(a int, b int) int {
	return a + b

}

func Plus3(a int, b int, c int) int {
	return a + b + c
}

// everything should call inside of main function
func main() {
	result1 := Plus2(2, 3)
	Hello()
	Plus1(1, 2)
	fmt.Println(result1)

	result3 := Plus3(2, 2, 6)
	fmt.Println(result3)

}
