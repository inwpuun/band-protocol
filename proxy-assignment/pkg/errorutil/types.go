package errorutil

type ErrorMeta[T any] struct {
	Type string
	Data T
}
