package main

import "fmt"

func main() {

	// index := 1

	// for index <= 10 {
	// 	fmt.Println(index)
	// 	index = index + 1
	// }

	// for index := 1; index <= 10; index++ {
	// 	fmt.Println(index)
	// }
	index := 0
	var names = [3]string{"Mehmet", "Can", "Doğan"}

	for index < len(names) {
		fmt.Println(names[index])
		index++
	}

}
