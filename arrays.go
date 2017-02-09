package main

import "fmt"

func main() {

	var prices [5]int
	fmt.Println("prices:", prices)

	prices[2] = 3
	fmt.Println("prices after update:", prices)

	ages := [5]int{21, 34, 23, 45}
	fmt.Println("ages:", ages)

}
