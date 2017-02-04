package main

import "fmt"

func main() {

	number := 333

	if number%2 == 0 {
		numberString := fmt.Sprintf("%d", number)
		fmt.Println(numberString + " is even")
	} else {
		numberString := fmt.Sprintf("%d", number)
		fmt.Println(numberString + " is odd")
	}

}
