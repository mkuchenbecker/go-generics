package main

type SliceIterator[T any] []T

func (s SliceIterator[T]) Next() (T, Iterator[T]) {
	return s[0], s[1:]
}

func (s SliceIterator[T]) HasNext() bool {
	return len(s) > 0
}

type Iterator[T any] interface {
	Next() (T, Iterator[T])
	HasNext() bool
}

type O any

// type Slice[T any] []T

// func (f Slice[T]) Map(mapper func(t T) O) Slice[O] {
// 	return Map[T, O](f, mapper)
// }

// func (f Slice[T]) Filter(filter func(t T) bool) Slice[T] {
// 	return Filter[T](f, filter)
// }

// func (f Slice[T]) Do(doer func(t T)) {
// 	Do[T](f, doer)
// }

func Map[T any, O any](it Iterator[T], f func(t T) O) Iterator[O] {
	r := make([]O, 0)

	for item, it := it.Next(); it.HasNext(); item, it = it.Next() {
		r = append(r, f(item))
	}
	return SliceIterator[O](r)
}

func Do[T any](it Iterator[T], doer func(t T)) {
	for item, it := it.Next(); it.HasNext(); item, it = it.Next() {
		doer(item)
	}
}

func Filter[T any](it Iterator[T], filter func(t T) bool) Iterator[T] {
	r := make([]T, 0)
	for item, it := it.Next(); it.HasNext(); item, it = it.Next() {
		if filter(item) {
			r = append(r, item)
		}
	}
	return SliceIterator[T](r)
}
