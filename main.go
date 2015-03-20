package main

import (
	"fmt"

	//	"github.com/gonum/blas/blas64"
	. "goDeep/nn"
)

func main() {
	//	v := blas64.Vector{Inc: 1, Data: []float64{1, 1, 1}}
	size := []int{10, 5, 4}
	nn := NewNN(size, 0.1, "sigmoid")
	input := []float64{0.1, 0.2, 0.3, 0.4, 0.5, 0.6, 0.7, 0.8, 0.9, 1}
	nn.SetInput(input)
	//fmt.Println(nn)
	nn.Forward()
	fmt.Println(nn)
	fmt.Println(nn.GetOutput())
	//fmt.Println("v has length:", blas64.Nrm2(len(v.Data), v))
}
