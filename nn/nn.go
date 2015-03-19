// nn
package nn

import (
	"fmt"
	"github.com/matrix/"
	"goDeep/utils"
)

type NeuralNetwork struct {
	HiddenLayer      []float64
	InputLayer       []float64
	OutputLayer      []float64
	WeightHidden     [][]float64
	WeightOutput     [][]float64
	ErrOutput        []float64
	ErrHidden        []float64
	LastChangeHidden [][]float64
	LastChangeOutput [][]float64
	Regression       bool
	Rate1            float64 //learning rate
	Rate2            float64
}

func main() {
	fmt.Println("Hello World!")
}
