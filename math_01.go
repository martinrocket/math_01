package main

import (
	"fmt"
	"math"
)

func main() {
	for i := -5; i <= 5; i++ {
		fmt.Print(i, " ")
		fmt.Println(math.Abs(float64(i)))
	}

}
