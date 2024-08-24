package interfaces

// Module to hold generic data types

type Num interface {
	int | int8 | int16 | int32 | int64 | float32 | float64
}
type NumArray interface {
	[]int | []int8 | []int16 | []int32 | []int64 | []float32 | []float64
}

func Index[T NumArray](arr T, index int) interface{} {
	switch v := any(arr).(type) {
	case []int:
		return v[index]
	case []int8:
		return v[index]
	case []int16:
		return v[index]
	case []int32:
		return v[index]
	case []int64:
		return v[index]
	case []float32:
		return v[index]
	case []float64:
		return v[index]
	default:
		panic("unsupported type")
	}
}

type NumMatrix interface {
	[][]int | [][]int8 | [][]int16 | [][]int32 | [][]int64 | [][]float32 | [][]float64
}
type NumArei interface {
	[]int | [][]int | []float64 | [][]float64
}
