package utils

func ToP[T comparable](p T) *T {
	return &p
}
