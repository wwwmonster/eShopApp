package helper

import (
	"iter"
	"slices"
)

type Iterator[V any] struct {
	seq iter.Seq[V]
}

func Stream[V any](src []V) *Iterator[V] {
	return &Iterator[V]{seq: slices.Values(src)}
}

func (it *Iterator[V]) Filter(pred func(V) bool) *Iterator[V] {
	return &Iterator[V]{
		seq: func(yield func(V) bool) {
			for v := range it.seq {
				if !pred(v) {
					continue
				}
				if !yield(v) {
					return
				}
			}
		},
	}
}

func (it *Iterator[V]) Collect() []V {
	return slices.Collect(it.seq)
}

func (it *Iterator[V]) ForEach(f func(V)) {
	for x := range it.seq {
		f(x)
	}
}

func (it *Iterator[V]) Map(f func(V) V) *Iterator[V] {
	return &Iterator[V]{
		seq: func(yield func(V) bool) {
			for v := range it.seq {
				if !yield(f(v)) {
					return
				}
			}
		},
	}
}
