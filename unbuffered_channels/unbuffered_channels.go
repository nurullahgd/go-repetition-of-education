package main

import "fmt"

func main() {

	myChannel := make(chan string)

	// go func() {
	// 	myChannel <- "Hello World" // This will block the go routine
	// }()
	// // This will block the main thread

	// message, isClosed := <-myChannel
	// println(message, isClosed)

	// myChannel <- "Hello World"

	go func() {
		message := <-myChannel
		println(message)
	}()
	go func() {
		myChannel <- "Hello World"
	}()

	fmt.Println("End of main")

}
