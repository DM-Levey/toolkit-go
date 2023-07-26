package slice

type List[T any] interface {
	Add(i int, T any) []T
	Delete(i int) []T
}
