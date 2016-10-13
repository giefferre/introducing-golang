package main

import (
	"fmt"
)

// START1 OMIT
func multiply(num1, num2 int) int {
	return num1 * num2
}

func main() {
	a := 7
	b := 5

	output := fmt.Sprintf(
		"Given a = %d, b = %d, the product is: %d",
		a,
		b,
		multiply(a, b),
	)
	fmt.Println(output)
}

//STOP1 OMIT
