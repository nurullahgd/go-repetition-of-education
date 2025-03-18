package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// go f1()
	// go f2()
	// fmt.Println("End of main")
	wg := sync.WaitGroup{}
	startTime := time.Now()
	wg.Add(2)
	go func() {
		defer wg.Done()
		fmt.Println("End of f1")
	}()
	go func() {
		defer wg.Done()
		fmt.Println("End of f2")
	}()
	wg.Wait() //Blockd

	fmt.Println("End of main", time.Now().Sub(startTime))
}

// func f1() {
// 	fmt.Println("End of f1")
// }

// func f2() {
// 	fmt.Println("End of f2")
// }
