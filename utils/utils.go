package utils

import "fmt"

// Dot product of two vectors is a sum of products of consecutive vector element. vectors must have same len
// a: inputs; b: weights
func DotProduct(inputs, weights []float64) float64 {
	if len(inputs) != len(weights) {
		return 0
	}
	var product float64
	for i := range inputs {
		product += inputs[i] * weights[i]
	}
	return product
}

func MatrixMultiply(matA, matB [][]float64) ([][]float64, error) {
	rowA, colA := len(matA), len(matA[0])
	rowB, colB := len(matB), len(matB[0])

	if colA != rowB {
		return [][]float64{}, fmt.Errorf("number of cols in matrix A must be equal to rows in matrix B")
	}
	result := make([][]float64, rowA)
	for i := range result {
		result[i] = make([]float64, colB)
	}
	//transpose matrix B for better caching. easier to transform the cols of matB to rows for DotProduct
	matB = TransposeMatrix(matB)

	// i: row index for matA & result; j: row for matB; k: col index for matA & matB
	for i := 0; i < rowA; i++ {
		for j := 0; j < colB; j++ {
			// var sum float64
			// for k := 0; k < colA; k++ {
			// 	sum += matA[i][k] * matB[j][k]
			// }
			result[i][j] = DotProduct(matA[i], matB[j])
		}
	}

	return result, nil
}

// rows to cols and cols to row
func TransposeMatrix(matrix [][]float64) [][]float64 {
	//make rows
	res := make([][]float64, len(matrix[0]))
	for i := range res {
		//make cols
		res[i] = make([]float64, len(matrix))
	}
	for i := 0; i < len(res); i++ {
		for j := 0; j < len(res[0]); j++ {
			res[i][j] = matrix[j][i]
		}
	}
	return res
}
