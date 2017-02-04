package main

import "fmt"

func main() {

	for order := 1; order <= 10; order++ {
		fmt.Println("Order n°", order)
	}

	for order := 1; order < 5; order++ {
		if order%2 == 0 {
			continue
		}
		fmt.Println(order)
	}

}
