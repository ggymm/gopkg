package gopkg

func If[T any](expr bool, v1, v2 T) T {
	if expr {
		return v1
	} else {
		return v2
	}
}
