package main

import (
	"fmt"
	"math"
)

func main() {
	h := 1.75
	var weight float64 = 90
	var IMT = weight / math.Pow(h, 2)
	fmt.Print(IMT)
}
