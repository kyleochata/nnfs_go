package main

import (
	"fmt"
	"nnfs_go/data"
	"nnfs_go/utils"

	"gonum.org/v1/gonum/mat"
)

type Neuron struct {
	weights []float64
	bias    float64
}

func RunModel() {
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

	a_raw := []float64{1, 2, 3}
	b_raw := []float64{2, 3, 4}

	b := utils.TransposeMatrix([][]float64{b_raw}) //[][]float64{{2}, {3}, {4}}
	a_mat := mat.NewDense(1, 3, a_raw)
	b_mat := mat.NewDense(3, 1, b_raw)
	//b_matT := b_mat.T()
	var result_mat mat.Dense
	result_mat.Mul(a_mat, b_mat)

	res, err := utils.MatrixMultiply([][]float64{a_raw}, b)
	if err != nil {
		fmt.Println("Error matrixMultiply", err)
	}

	fmt.Println("bmat using mat", mat.Formatted(b_mat))
	fmt.Println("res mat", mat.Formatted(&result_mat))
	fmt.Println("res matrixMult", res)

	b_inputs := [][]float64{
		inputs,
		{2.0, 5.0, -1.0, 2.0},
		{-1.5, 2.7, 3.3, -0.8},
	}

	weights := [][]float64{weights1, weights2, weights3}
	biases := []float64{2.0, 3.0, 0.5}

	fmt.Println(utils.TransposeMatrix(weights))

	layerout_1, err := utils.MatrixMultiply(b_inputs, utils.TransposeMatrix(weights))
	if err != nil {
		fmt.Println("matrixMult batch inputs err", err)
		return
	}

	for i := range layerout_1 {
		for j := 0; j < len(biases); j++ {
			layerout_1[i][j] += biases[j]
		}
	}
	fmt.Println("==== layerout_1 ====")
	for i, row := range layerout_1 {
		fmt.Printf("i: %v row: ", i)
		for _, val := range row {
			fmt.Printf("%.3f ", val) // Format to 2 decimal places
		}
		fmt.Println() // New line after each row
	}
	fmt.Println("==== END LAYER 1 ====")
	second_weights := [][]float64{
		{0.1, -0.14, 0.5},
		{-0.5, 0.12, -0.33},
		{-0.44, 0.73, -0.13},
	}
	second_biases := []float64{-1, 2, -0.5}

	layerout_2, err := utils.MatrixMultiply(layerout_1, utils.TransposeMatrix(second_weights))
	if err != nil {
		fmt.Println("matrixMult batch inputs err", err)
		return
	}

	for i := range layerout_2 {
		for j := 0; j < len(biases); j++ {
			layerout_2[i][j] += second_biases[j]
		}
	}
	fmt.Println("==== layerout_2 ====")
	for i, row := range layerout_2 {
		fmt.Printf("i: %v row: ", i)
		for _, val := range row {
			fmt.Printf("%.3f ", val) // Format to 2 decimal places
		}
		fmt.Println() // New line after each row
	}
	fmt.Println("==== END LAYER 2 ====")
}

func main() {
	RunModel()
	x, y := data.SpiralData(100, 3)

	fmt.Println(x)
	fmt.Println(y)
}
