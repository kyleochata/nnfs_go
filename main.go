package main

import (
	"fmt"
	"nnfs_go/utils"
)

type Neuron struct {
	weights []float64
	bias    float64
}

func main() {
	inputs := []float64{1, 2, 3, 2.5}

	weights1 := []float64{0.2, 0.8, -0.5, 1.0}
	weights2 := []float64{0.5, -0.91, 0.26, -0.5}
	weights3 := []float64{-0.26, -0.27, 0.17, 0.87}
	neuron_1 := Neuron{weights1, 2.0}
	neuron_2 := Neuron{weights2, 3.0}
	neuron_3 := Neuron{weights3, 0.5}
	fmt.Printf("neuron1: %v\n neuron2: %v\n neuron3: %v\n", neuron_1, neuron_2, neuron_3)

	outArr := [3]float64{0, 0, 0}
	neuronCluster := []Neuron{neuron_1, neuron_2, neuron_3}

	for i := range outArr {
		outArr[i] = utils.DotProduct(inputs, neuronCluster[i].weights) + neuronCluster[i].bias
	}

	fmt.Println("Outputs", outArr)

	a := [][]float64{{1, 2, 3}}
	b := utils.TransposeMatrix([][]float64{{2, 3, 4}}) //[][]float64{{2}, {3}, {4}}

	res, err := utils.MatrixMultiply(a, b)
	if err != nil {
		fmt.Println("Error matrixMultiply", err)
	}
	fmt.Println("res matrixMult", res)

	b_inputs := [][]float64{
		inputs,
		{2.0, 5.0, -1.0, 2.0},
		{-1.5, 2.7, 3.3, -0.8},
	}

	weights := [][]float64{weights1, weights2, weights3}
	biases := []float64{2.0, 3.0, 0.5}

	fmt.Println(utils.TransposeMatrix(weights))

	layer_outs, err := utils.MatrixMultiply(b_inputs, utils.TransposeMatrix(weights))
	if err != nil {
		fmt.Println("matrixMult batch inputs err", err)
		return
	}

	for i := range layer_outs {
		for j := 0; j < len(biases); j++ {
			layer_outs[i][j] += biases[j]
		}
	}

	for i, row := range layer_outs {
		fmt.Printf("i: %v row: ", i)
		for _, val := range row {
			fmt.Printf("%.3f ", val) // Format to 2 decimal places
		}
		fmt.Println() // New line after each row
	}
}
