package nn

import (
	"fmt"
	. "github.com/skelterjohn/go.matrix"
	"math"
	"math/rand"
)

type NN struct {
	layers       []*DenseMatrix
	weights      []*DenseMatrix
	learningRate float64
}

func initWeight(m, n int) *DenseMatrix {

	matrixArray := make([]float64, (m+1)*n)
	//matrix := make(Weight, m+1)
	for i := 0; i < m+1; i++ {
		//matrix[i] = make([]float64, n)
		for j := 0; j < n; j++ {
			matrixArray[i*n+j] = (rand.Float64() - 0.5) * 2 * math.Sqrt(float64(6.0/(m+1+n)))
		}
	}
	matrix := MakeDenseMatrix(matrixArray, m+1, n)
	return matrix
}

func (this *NN) GetOutput() []float64 {

	return this.layers[len(this.layers)-1].Array()
}

func NewNN(size []int, learningRate float64) (nn *NN) {
	network := NN{}
	levelNum := len(size)
	network.layers = make([]*DenseMatrix, levelNum)
	network.weights = make([]*DenseMatrix, levelNum-1)
	network.learningRate = learningRate
	for i := 0; i < levelNum; i++ {
		layer := make([]float64, size[i])
		network.layers[i] = MakeDenseMatrix(layer, 1, size[i])
		if i != levelNum-1 {
			network.weights[i] = initWeight(size[i], size[i+1])
		}
	}
	return &network
}

func (this *NN) SetInput(input []float64) {
	this.layers[0] = MakeDenseMatrix(input, 1, len(input))
}

func (this *NN) String() string {
	s := "weight:"
	for i := 0; i < len(this.weights); i++ {
		s += fmt.Sprintf("%d", i) + ":" + String(this.weights[i])
	}
	return s
}

func (this *NN) Forward() {
	weightNum := len(this.weights)
	for i := 0; i < weightNum; i++ {
		rowVector := this.layers[i].Array()
		biasedLayer := MakeDenseMatrix(append(rowVector, 1.0), 1, len(rowVector)+1)
		matrix, err := biasedLayer.TimesDense(this.weights[i])

		if err != nil {
			this.layers[i+1] = matrix
		}
		//this.layers[i+1].MulVec(this.weights[i], true, biasedLayer)

	}
}
