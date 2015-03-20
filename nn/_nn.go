// nn
package nn

import (
	"fmt"
	. "github.com/gonum/matrix/mat64"
	"math"
	"math/rand"
	//"goDeep/utils"
)

//type NeuralNetwork struct {
//	HiddenLayer      []float64
//	InputLayer       []float64
//	OutputLayer      []float64
//	WeightHidden     [][]float64
//	WeightOutput     [][]float64
//	ErrOutput        []float64
//	ErrHidden        []float64
//	LastChangeHidden [][]float64
//	LastChangeOutput [][]float64
//	Regression       bool
//	Rate1            float64 //learning rate
//	Rate2            float64
//}

//type Matrix Dense
type Weight struct {
	Dense
}

type NN struct {
	layers       []*Vector
	weights      []*Dense
	learningRate float64
}

func initWeight(m, n int) *Dense {

	matrixArray := make([]float64, (m+1)*n)
	//matrix := make(Weight, m+1)
	for i := 0; i < m+1; i++ {
		//matrix[i] = make([]float64, n)
		for j := 0; j < n; j++ {
			matrixArray[i*n+j] = (rand.Float64() - 0.5) * 2 * math.Sqrt(float64(6.0/(m+1+n)))
		}
	}
	matrix := NewDense(m+1, n, matrixArray)
	return matrix
}

func (this *NN) GetOutput() Vector {
	return *this.layers[len(this.layers)-1]
}

func NewNN(size []int, learningRate float64) (nn *NN) {
	network := NN{}
	levelNum := len(size)
	network.layers = make([]*Vector, levelNum)
	network.weights = make([]*Dense, levelNum-1)
	network.learningRate = learningRate
	for i := 0; i < levelNum; i++ {
		layer := make([]float64, size[i])
		network.layers[i] = NewVector(size[i], layer)
		if i != levelNum-1 {
			network.weights[i] = initWeight(size[i], size[i+1])
		}
	}
	return &network
}

func (this *NN) SetInput(input []float64) {
	this.layers[0] = NewVector(len(input), input)
}

func (this *NN) Forward() {
	weightNum := len(this.weights)
	for i := 0; i < weightNum; i++ {
		biasedLayer := NewVector(this.layers[i].Len()+1, append(this.layers[i].RawVector().Data, 1.0))
		this.layers[i+1].MulVec(this.weights[i], true, biasedLayer)

	}
}

func main() {
	fmt.Println("Hello World!")
}
