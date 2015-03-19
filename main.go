package main

import (
	"fmt"

	"github.com/gonum/blas/blas64"
)

func main() {
	v := blas64.Vector{Inc: 1, Data: []float64{1, 1, 1}}
	fmt.Println("v has length:", blas64.Nrm2(len(v.Data), v))
}
