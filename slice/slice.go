package main

import "fmt"

func main() {
	var courseName []string
	courseName = []string{"Java", "Go"}
	fmt.Println(courseName)
	courseName = append(courseName, "Python", "C", "React", "C#", "TS")
	fmt.Println(courseName)

	courseWeb := courseName[4:7]
	fmt.Println(courseWeb)

	courseWeb = courseName[:4]
	fmt.Println(courseWeb)
}
