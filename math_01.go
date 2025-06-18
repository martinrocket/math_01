package main

import (
	"fmt"
	"math"
	"strings"
	
)

func absandsqrt(x int) (float64, float64, float64) {
	for i:= 2; i < 4; i++ {
		return math.Pow(math.Abs(float64(x)), i)
	}
}

for i := 0; i < 5; i++ {
        fmt.Println(i) // Prints 0, 1, 2, 3, 4
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

