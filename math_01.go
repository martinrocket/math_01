package main

import (
	"fmt"
	"math"
	"strings"
	
)

func absandsqrt(x int) (float64, float64, float64) {
	var a, b, c float64
	a = math.Pow(math.Abs(float64(x)), 2)
	b = math.Pow(math.Abs(float64(x)), 3)
	c = math.Pow(math.Abs(float64(x)), 4)
	return a, b, c
    }





func main() {
	/* Starting out with a loop that goes from neg 5 to
	positive 5 then does a math.Abs (absolute value)
	calculation to i*/

	var myList string
	myList = "abcdefg"
	fmt.Println(myList)

	for i := -5; i <= 5; i++ {
		fmt.Print(i, " ")
		fmt.Print(math.Abs(float64(i)), " ")
		fmt.Println(absandsqrt(i))
	}
	fmt.Print("\n\n")
	var myString strings.Builder
	fmt.Fprintf(&myString, "abcd")
	fmt.Println(myString.String())
	

}

