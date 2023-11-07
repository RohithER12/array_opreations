package arrayoprations

func SumInt(array []int) int {
	var sum int
	for _, a := range array {
		sum += a
	}
	return sum
}

func SumFloat64(array []float64) float64 {
	var sum float64
	for _, a := range array {
		sum += a
	}
	return sum
}

func SumFloat32(array []float32) float32 {
	var sum float32
	for _, a := range array {
		sum += a
	}
	return sum
}

func SumInt32(array []int32) int32 {
	var sum int32
	for _, a := range array {
		sum += a
	}
	return sum
}

func AverageInt(array []int) int {
	var sum int
	for _, a := range array {
		sum += a
	}
	return sum / len(array)
}
