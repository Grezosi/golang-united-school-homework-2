package main

import (
	"fmt"
	"math"
)

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3),SidesTriangle(==3), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

func CalcSquare(sideLen float64, sidesNum int64) float64 {
	const (
		SidesTriangle = 3
		SidesSquare   = 4
	)
	switch {
	case sidesNum == SidesTriangle:
		return SidesTriangle * sideLen
	case sidesNum == SidesSquare:
		return SidesSquare * sideLen
	case sidesNum == 0:
		return sideLen * sideLen * math.Pi
	}
	return 0
}
func main() {
	fmt.Println(CalcSquare(2, 0))
}