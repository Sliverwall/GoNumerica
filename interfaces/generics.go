package interfaces

type Num interface {
	int | int8 | int16 | int32 | int64 | float32 | float64
}
type NumArray interface {
	[]int | []int8 | []int16 | []int32 | []int64 | []float32 | []float64
}
type NumMatrix interface {
	[][]int | [][]int8 | [][]int16 | [][]int32 | [][]int64 | [][]float32 | [][]float64
}
