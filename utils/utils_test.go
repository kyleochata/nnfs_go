package utils

import (
	"reflect"
	"testing"
)

type MatrixMultiplyTestCase struct {
	name    string
	A       [][]float64
	B       [][]float64
	want    [][]float64
	wantErr bool
}

var matrixTestCases = []MatrixMultiplyTestCase{
	{
		name: "valid 2x3 * 3x2 multiplication",
		A: [][]float64{
			{1, 2, 3},
			{4, 5, 6},
		},
		B: [][]float64{
			{7, 8},
			{9, 10},
			{11, 12},
		},
		want: [][]float64{
			{58, 64},
			{139, 154},
		},
		wantErr: false,
	},
	{
		name:    "incompatible dimensions",
		A:       [][]float64{{1, 2}},
		B:       [][]float64{{1}},
		want:    nil,
		wantErr: true,
	},
	{
		name:    "1x3 x 3x1",
		A:       [][]float64{{1, 2, 4}},
		B:       [][]float64{{2}, {3}, {4}},
		want:    [][]float64{{20}},
		wantErr: false,
	},
}

func TestMatrixMultiply(t *testing.T) {
	for _, testCase := range matrixTestCases {
		t.Run(testCase.name, func(t *testing.T) {
			got, err := MatrixMultiply(testCase.A, testCase.B)
			if (err != nil) != testCase.wantErr {
				t.Errorf("MatrixMultiply() error = %v, wantErr = %v", err, testCase.wantErr)
				return
			}
			if !testCase.wantErr && !reflect.DeepEqual(got, testCase.want) {
				t.Errorf("MatrixMultiply() = %v, want = %v", got, testCase.want)
			}
		})
	}
}

type TransposeMatrixCase struct {
	name string
	A    [][]float64
	want [][]float64
}

var TansMCases = []TransposeMatrixCase{
	{
		name: "transpose 2x1",
		A: [][]float64{
			{2},
			{1},
		},
		want: [][]float64{
			{2, 1},
		},
	},
	{
		name: "transpose 1x3",
		A: [][]float64{
			{1, 2, 3},
		},
		want: [][]float64{
			{1},
			{2},
			{3},
		},
	},
}

func TestTransposeMatrix(t *testing.T) {
	for _, testCase := range TansMCases {
		t.Run(testCase.name, func(t *testing.T) {
			got := TransposeMatrix(testCase.A)
			if !reflect.DeepEqual(got, testCase.want) {
				t.Errorf("TransposeMatrix() = %v, want = %v", got, testCase.want)
			}
		})
	}
}
