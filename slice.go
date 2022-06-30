package fn

import (
	"sync"

	"golang.org/x/exp/constraints"
)

// Map applies the function f to each element of the collection and returns a new collection with the results.
func Map[T any](xs []T, f func(T) T) []T {
	ys := make([]T, len(xs))
	for i, x := range xs {
		ys[i] = f(x)
	}
	return ys
}

// PMap applies the function f to each element of the collection in parallel and returns a new collection with the results.
func PMap[T any](xs []T, f func(T) T) []T {
	ys := make([]T, len(xs))

	wg := new(sync.WaitGroup)

	for i := range xs {
		wg.Add(1)
		go func(i int) {
			ys[i] = f(xs[i])

			wg.Done()
		}(i)
	}

	wg.Wait()

	return ys
}

// Reduce applies the function f to each element of the collection and returns the result.
func Reduce[T any, R any](acc R, xs []T, f func(R, T) R) R {
	for _, x := range xs {
		acc = f(acc, x)
	}
	return acc
}

// ForEach applies the function f to each element of the collection.
func ForEach[T any](xs []T, f func(T)) {
	for _, x := range xs {
		f(x)
	}
}

func PForEach[T any](xs []T, f func(T)) {
	wg := new(sync.WaitGroup)

	for _, x := range xs {
		wg.Add(1)
		go func(x T) {
			f(x)

			wg.Done()
		}(x)
	}

	wg.Wait()
}

func Filter[T any](xs []T, f func(T) bool) []T {
	ys := make([]T, 0, len(xs))
	for _, x := range xs {
		if f(x) {
			ys = append(ys, x)
		}
	}
	return ys
}

func Find[T any](xs []T, f func(T) bool) (T, bool) {
	for _, x := range xs {
		if f(x) {
			return x, true
		}
	}
	var d T
	return d, false
}

func Zip[T any, U any](xs []T, ys []U) []Tuple[T, U] {
	zs := make([]Tuple[T, U], len(xs))
	for i := range xs {
		zs[i] = Tuple[T, U]{xs[i], ys[i]}
	}
	return zs
}

func Reverse[T any](xs []T) []T {
	ys := make([]T, len(xs))
	for i, x := range xs {
		ys[len(xs)-i-1] = x
	}
	return ys
}

func Max[T constraints.Ordered](xs []T) T {
	m := xs[0]
	for i := 1; i < len(xs); i++ {
		if xs[i] > m {
			m = xs[i]
		}
	}
	return m
}

func Min[T constraints.Ordered](xs []T) T {
	m := xs[0]
	for i := 1; i < len(xs); i++ {
		if xs[i] < m {
			m = xs[i]
		}
	}
	return m
}

func GroupBy[K comparable, V comparable](m map[K]V) map[V][]K {
	res := make(map[V][]K)

	for k, v := range m {
		res[v] = append(res[v], k)
	}

	return res
}

func GroupByFn[K comparable, V any, R comparable](m map[K]V, fn func(V) R) map[R][]K {
	res := make(map[R][]K)

	for k, v := range m {
		key := fn(v)
		res[key] = append(res[key], k)
	}

	return res
}
