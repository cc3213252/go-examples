package main

import (
	"fmt"
	"math"
)

type Ver struct {
	X, Y float64
}

func (v Ver) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Ver) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Ver{3, 4}
	v.Scale(2)
	fmt.Println("result:", v.Abs())
}
