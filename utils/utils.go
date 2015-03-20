package utils

import (
	"math"
	//	"math/rand"
)

type ActivateFunc func(float64) float64

func Activate(activate string) ActivateFunc {
	switch activate {
	case "sigmoid":
		return Sigmoid
	default:
		return Sigmoid
	}

}

func Sigmoid(x float64) float64 {
	return 1.0 / (1.0 + math.Exp(-x))
}

func Dsigmoid(x float64) float64 {
	return x * (1.0 - x)
}
