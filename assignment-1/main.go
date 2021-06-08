package main

import "fmt"

func main() {
	// define a empty slice of int
	numbers := []int{}

	// for loop to fill the slice of int with 0-10
	for i := 0; i <= 10; i++ {
		numbers = append(numbers, i)
	}

	// loop throw the slice of int mod the number by 2
	// if % 2 is 0 then the number is even
	// else the number is odd
	for _, number := range numbers {
		if number%2 == 0 {
			fmt.Println(number, "is even")
		} else {
			fmt.Println(number, "is odd")
		}
	}
}
