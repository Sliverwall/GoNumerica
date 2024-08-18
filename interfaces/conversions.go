package interfaces

// Module for handling data type conversions

// ConvertIntArrToFloat64Arr takes a given []int and returns the array casted a []float64
func ConvertIntArrToFloat64Arr(intArr []int) []float64 {
	// Make float64 slice, set hint = to given intArr length
	floatArr := make([]float64, len(intArr))
	// Iterate through each element in intArray then cast as float64 and set to floatArr
	for element := range floatArr {
		floatArr[element] = float64(intArr[element])
	}
	// Return floatArr after full iteration
	return floatArr
}

// ConvertFloat64ArrToFloat64Arr takes a given []float64 and returns the array casted a []int
func ConvertFloat64ArrToInt64Arr(float64Arr []float64) []int {
	// Make int slice, set hint = to given float64Arr length
	intArr := make([]int, len(float64Arr))
	// Iterate through each element in float64Array then cast as int and set to intArr
	for element := range intArr {
		intArr[element] = int(float64Arr[element])
	}
	// Return intArr after full iteration
	return intArr
}
