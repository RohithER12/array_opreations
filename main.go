package main

import (
	"fmt"

	arrayoprations "github.com/RohithER12/array_opreations/arrayoprations/array_oprations"
)

func main() {
	intArray := []int{1, 2, 3, 4, 5}
	float64Array := []float64{1.1, 2.2, 3.3, 4.4, 5.5}
	float32Array := []float32{1.1, 2.2, 3.3, 4.4, 5.5}
	int32Array := []int32{1, 2, 3, 4, 5}

	sumInt := arrayoprations.SumInt(intArray)
	sumFloat64 := arrayoprations.SumFloat64(float64Array)
	sumFloat32 := arrayoprations.SumFloat32(float32Array)
	sumInt32 := arrayoprations.SumInt32(int32Array)
	averageInt := arrayoprations.AverageInt(intArray)

	fmt.Printf("Sum of intArray: %d\n", sumInt)
	fmt.Printf("Sum of float64Array: %f\n", sumFloat64)
	fmt.Printf("Sum of float32Array: %f\n", sumFloat32)
	fmt.Printf("Sum of int32Array: %d\n", sumInt32)
	fmt.Printf("Average of intArray: %d\n", averageInt)
}
