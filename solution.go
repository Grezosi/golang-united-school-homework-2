package square

import "math"

type SidesNumber int64

const (
	SidesTriangle SidesNumber = 3
	SidesSquare   SidesNumber = 4
	SidesCircle   SidesNumber = 0
)

func CalcSquare(sideLen float64, sidesNum SidesNumber) float64 {

	switch sidesNum {
	case SidesTriangle:
		return math.Pow(sideLen, 3)
	case SidesSquare:
		return math.Pow(sideLen, 2)
	case SidesCircle:
		return sideLen * sideLen * math.Pi
	default:
		return 0

	}
}
