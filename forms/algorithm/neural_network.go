package algorithm

import (
	deep "github.com/patrikeh/go-deep"
	"math/rand"
	"runtime"
)

func Nn_algo(inputx[][]float64, result, prueba []float64) (float64) {
	rand.Seed(0)
	n := deep.NewNeural(&deep.Config{
		Inputs:     14,
		Layout:     []int{3, 3, 1},
		Activation: deep.ActivationSigmoid,
		Mode:       deep.ModeBinary,
		Weight:     deep.NewUniform(.25, 0),
		Bias:       true,
	})

	var solving Examples

	for i:= 0; i < len(result); i++ {
		solving = append(solving, Examples{{[]float64{inputx[i][0], inputx[i][1],inputx[i][2],inputx[i][3],inputx[i][4],inputx[i][5],inputx[i][6],inputx[i][7],inputx[i][8],inputx[i][9],inputx[i][10],inputx[i][11],inputx[i][12],inputx[i][13]}, []float64{result[i]}}}...)
	}

	const iterations = 4000
	solver := NewAdam(0.01, 0.9, 0.999, 1e-8)
	trainer := NewBatchTrainer(solver, iterations, len(solving)/2, runtime.NumCPU())
	trainer.Train(n, solving, solving, iterations)
	return float64(n.Predict(prueba)[0])
}