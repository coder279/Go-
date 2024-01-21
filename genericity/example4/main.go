package main

type PredeclaredSingendInteger interface {
	int | int8 | int64 | int16 | int32
}

type SignedInteger interface {
	~int | ~int8 | ~int64 | ~int16 | ~int32
}

func GreaterThan[T PredeclaredSingendInteger](a, b T) bool {
	if a > b {
		return true
	}
	return false
}
func main() {
	var a int = 10
	var b int = 200
	GreaterThan(a, b)
}
