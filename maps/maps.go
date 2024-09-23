package main

import "fmt"

var product = make(map[string]float64)

func main() {
	fmt.Println("Product =", product)

	// add
	product["Macbook"] = 40000
	product["Iphone"] = 30000
	product["Ipad"] = 20000

	fmt.Println("Product =", product)

	// remove a data in maps
	delete(product, "Ipad")

	fmt.Println("Product =", product)

	// Update data in maps
	product["Iphone"] = 35000
	fmt.Println("Product =", product)

	// access value in maps data

	value1 := product["Iphone"]
	fmt.Println(value1)

	// Maps
	courseName := map[string]string{"101": "Java", "102": "Go"}
	fmt.Println(courseName)

}
