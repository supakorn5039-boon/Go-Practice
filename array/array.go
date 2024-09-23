package main

import "fmt"

// init array
var productName [4]string

func main() {
	price := [4]float32{40000, 30000, 20000, 2000}
	productName[0] = "Macbook"
	productName[1] = "Ipad"
	productName[2] = "IPhone"
	productName[3] = "AirPods"
	fmt.Println(productName)
	fmt.Println(price)
}
