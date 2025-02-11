package main

import "fmt"

func main() {
	//Fixed
	//var names [3]string
	// names[0] = "Mehmet"
	// names[1] = "Can"
	// names[2] = "Doğan"

	// var names = [4]string{"Mehmet", "Can", "Doğan","Gündoğdu"}

	// fmt.Println(names[0:2])

	var names = []string{"Mehmet", "Can", "Doğan"}
	names = append(names, "Seks")
	fmt.Println(names)
}
