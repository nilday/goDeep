// nn
package nn

import (
	"fmt"
	. "github.com/gonum/matrix/mat64"
	u "goDeep/utils"
	"math"
	"math/rand"
)

type Weight struct {
	Dense
}

type NN struct {
	layers       []*Vector
	weights      []*Dense
	learningRate float64
	activateFunc u.ActivateFunc
}

func initWeight(m, n int) *Dense {

	matrixArray := make([]float64, (m+1)*n)
	for i := 0; i < m+1; i++ {
		for j := 0; j < n; j++ {
			faninfanout := math.Sqrt(float64(6.0) / float64(m+1+n))

			matrixArray[i*n+j] = (rand.Float64() - 0.5) * 2 * faninfanout
		}
	}
	matrix := NewDense(m+1, n, matrixArray)
	return matrix
}

func (this *NN) String() string {
	s := "weight:"
	i := 0
	for ; i < len(this.weights); i++ {
		s += fmt.Sprintf("layer %d", i) + ":\n" + fmt.Sprint(this.layers[i]) + "\n"
		s += fmt.Sprintf("weight %d", i) + ":\n" + fmt.Sprint(this.weights[i]) + "\n"
	}

	s += fmt.Sprintf("layer %d", i) + ":\n" + fmt.Sprint(this.layers[i]) + "\n"

	return s
}

func (this *NN) GetOutput() Vector {
	return *this.layers[len(this.layers)-1]
}

func NewNN(size []int, learningRate float64, activate string) (nn *NN) {
	network := NN{}
	levelNum := len(size)
	network.layers = make([]*Vector, levelNum)
	network.weights = make([]*Dense, levelNum-1)
	network.learningRate = learningRate
	network.activateFunc = u.Activate(activate)
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
		//fmt.Println(this.layers[i+1])
		this.activate(this.layers[i+1])
		//fmt.Println(this.layers[i+1])
	}
}

func (this *NN) activate(layer *Vector) {
	for i := 0; i < len(layer.RawVector().Data); i++ {
		layer.RawVector().Data[i] = this.activateFunc(layer.RawVector().Data[i])
	}
}

func main() {
	fmt.Println("Hello World!")
}
