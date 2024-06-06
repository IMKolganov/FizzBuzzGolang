package main

import (
	"fmt"
)

func main() {
	fmt.Println("FizzBuzz app")

	for i := 1; i <= 100; i++ {
		var result = ""
		found := false

		if i%3 == 0 {
			result = "Fizz"
			found = true
		}
		if i%5 == 0 {
			result += "Buzz"
			found = true
		}

		if !found {
			fmt.Println(i)
		} else {
			fmt.Println(result)
		}
	}

}
