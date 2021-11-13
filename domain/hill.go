package models

import "math"

type Hill struct {
	Length float64
	Slope  float64
}

func (h Hill) GetHeight() float64 {
	return h.Length * h.Slope / (math.Sqrt(h.Slope*h.Slope + 1))
}
