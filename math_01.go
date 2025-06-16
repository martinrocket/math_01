package main

import (
	"fmt"
	"math"
)

func Absandsqrt(x int) float64 {
	return math.Pow(math.Abs(float64(x)), 2)
}

func main() {
	/* Starting out with a loop that goes from neg 5 to
	positive 5 then does a math.Abs (absolute value)
	calculation to i*/

	for i := -5; i <= 5; i++ {
		fmt.Print(i, " ")
		fmt.Print(math.Abs(float64(i)), " ")
		fmt.Println(Absandsqrt(i))
	}

}
