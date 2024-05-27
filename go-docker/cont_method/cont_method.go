package main

import (
	"fmt"
	"math"
)

type Float float64 //Floatの部分がabsメソッドの戻り値

func (f Float) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	f := Float(-math.Sqrt2)
	fmt.Println(f.Abs())
}
