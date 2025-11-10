package program

import (
	"iter"
	"slices"
)

func Stream[T any](it []T) iter.Seq[T] {
	return slices.Values(it)
}

func Filter[T any](it iter.Seq[T], keep func(T) bool) iter.Seq[T] {
	return func(yield func(T) bool) {
		for ele := range it {
			if keep(ele) {
				if !yield(ele) {
					return
				}
			}
		}
	}
}

func Map[T any](it iter.Seq[T], keep func(T) T) iter.Seq[T] {
	return func(yield func(T) bool) {
		for ele := range it {
			if !yield(keep(ele)) {
				return
			}
		}
	}
}

func Reduce[T any](it iter.Seq[T], keep func(T, T) T) iter.Seq[T] {
	return func(yield func(T) bool) {
		var num T
		count := 0
		for ele := range it {
			if count == 0 {
				num = ele
			} else {
				num = keep(num, ele)
			}
			count++
		}
		if !yield(num) {
			return
		}
	}
}
