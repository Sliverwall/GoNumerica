package interfaces

// Module to hold generic data types

type Num interface {
	int | int8 | int16 | int32 | int64 | float32 | float64
}
type NumArray interface {
	[]int | []int8 | []int16 | []int32 | []int64 | []float32 | []float64
}
type NumMatrix interface {
	[][]int | [][]int8 | [][]int16 | [][]int32 | [][]int64 | [][]float32 | [][]float64
}
type NumArei interface {
	[]int | [][]int | []float64 | [][]float64
}
