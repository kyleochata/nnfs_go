package data

import (
	"math"
	"math/rand"
)

// samples = number of points per spiral; classes = number of spirals
// return a slice of [x,y] coordinates and a slice that denotes what class the corresponding index's point
// belongs to
func SpiralData(samples int, classes int) ([][]float64, []uint8) {
	x := make([][]float64, samples*classes)
	y := make([]uint8, samples*classes)

	for i := 0; i < len(x); i++ {
		x[i] = make([]float64, 2)
	}

	for classNumber := 0; classNumber < classes; classNumber++ {
		rad := linespace(0, 1.0, samples)
		theta := linespace(float64(classNumber)*4, float64(classNumber+1)*4, samples)

		for i := range theta {
			theta[i] += 0.2 * rand.NormFloat64()
		}
		//convert polar to cartesian coords
		for ix := classNumber * samples; ix < samples*(classNumber+1); ix++ {
			idx := ix - classNumber*samples
			angle := theta[idx] * 2.5 // controls how tightly wound spiral is
			x[ix][0] = rad[idx] * math.Sin(angle)
			x[ix][1] = rad[idx] * math.Cos(angle)
			y[ix] = uint8(classNumber)
		}

	}
	return x, y
}

// linespace creates a slice of n values between start and end
func linespace(start, end float64, n int) []float64 {
	if n <= 1 {
		return []float64{start}
	}

	step := (end - start) / float64(n-1)
	result := make([]float64, n)
	for i := range n {
		result[i] = start + float64(i)*step
	}
	return result
}
