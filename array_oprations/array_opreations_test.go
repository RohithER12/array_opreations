package arrayoprations_test

import (
	"testing"

	arrayoprations "github.com/RohithER12/array_oprations"
)

func TestSumInt(t *testing.T) {
	testCases := []struct {
		input    []int
		expected int
	}{
		{[]int{1, 2, 3}, 6},
		{[]int{5, 5, 5, 5}, 20},
		{[]int{}, 0},
	}

	for _, tc := range testCases {
		result := arrayoprations.SumInt(tc.input)
		if result != tc.expected {
			t.Errorf("SumInt(%v) = %d; expected %d", tc.input, result, tc.expected)
		}
	}
}

func TestSumFloat64(t *testing.T) {
	testCases := []struct {
		input    []float64
		expected float64
	}{
		{[]float64{1.1, 2.2, 3.3}, 6.6},
		{[]float64{5.5, 5.5, 5.5, 5.5}, 22.0},
		{[]float64{}, 0.0},
	}

	for _, tc := range testCases {
		result := arrayoprations.SumFloat64(tc.input)
		if result != tc.expected {
			t.Errorf("SumFloat64(%v) = %f; expected %f", tc.input, result, tc.expected)
		}
	}
}

func TestSumFloat32(t *testing.T) {
	testCases := []struct {
		input    []float32
		expected float32
	}{
		{[]float32{1.1, 2.2, 3.3}, 6.6},
		{[]float32{5.5, 5.5, 5.5, 5.5}, 22.0},
		{[]float32{}, 0.0},
	}

	for _, tc := range testCases {
		result := arrayoprations.SumFloat32(tc.input)
		if result != tc.expected {
			t.Errorf("SumFloat32(%v) = %f; expected %f", tc.input, result, tc.expected)
		}
	}
}

func TestSumInt32(t *testing.T) {
	testCases := []struct {
		input    []int32
		expected int32
	}{
		{[]int32{1, 2, 3}, 6},
		{[]int32{5, 5, 5, 5}, 20},
		{[]int32{}, 0},
	}

	for _, tc := range testCases {
		result := arrayoprations.SumInt32(tc.input)
		if result != tc.expected {
			t.Errorf("SumInt32(%v) = %d; expected %d", tc.input, result, tc.expected)
		}
	}
}

func TestAverageInt(t *testing.T) {
	testCases := []struct {
		input    []int
		expected int
	}{
		{[]int{1, 2, 3}, 2},
		{[]int{5, 5, 5, 5}, 5},
		{[]int{}, 0},
	}

	for _, tc := range testCases {
		result := arrayoprations.AverageInt(tc.input)
		if result != tc.expected {
			t.Errorf("AverageInt(%v) = %d; expected %d", tc.input, result, tc.expected)
		}
	}
}
