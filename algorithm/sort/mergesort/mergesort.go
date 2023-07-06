package main

import (
	"sync"

	"github.com/TheAlgorithms/Go/math/min"

	"golang.org/x/exp/constraints"
)

// https://gmlwjd9405.github.io/2018/05/08/algorithm-merge-sort.html
func merge[T constraints.Ordered](a []T, b []T) []T {
	r := make([]T, len(a)+len(b))
	i := 0
	j := 0

	for i < len(a) && j < len(b) {
		if a[i] <= b[j] {
			r[i+j] = a[i]
			i++
		} else {
			r[i+j] = b[j]
			j++
		}
	}

	for i < len(a) {
		r[i+j] = a[i]
		i++
	}
	for j < len(b) {
		r[i+j] = b[j]
		j++
	}

	return r
}

// Merge Perform merge sort on a slice
func Merge[T constraints.Ordered](items []T) []T {
	if len(items) < 2 {
		return items
	}

	middle := len(items) / 2
	a := Merge(items[:middle])
	b := Merge(items[middle:])
	return merge(a, b)
}

func MergeIter[T constraints.Ordered](items []T) []T {
	for step := 1; step < len(items); step += step {
		for i := 0; i+step < len(items); i += 2 * step {
			tmp := merge(items[i:i+step], items[i+step:min.Int(i+2*step, len(items))])
			copy(items[i:], tmp)
		}
	}
	return items
}

// ParallelMerge Perform merge sort on a slice using goroutines
func ParallelMerge[T constraints.Ordered](items []T) []T {
	if len(items) < 2 {
		return items
	}

	if len(items) < 2048 {
		return Merge(items)
	}

	var wg sync.WaitGroup
	wg.Add(1)

	middle := len(items) / 2
	var a []T
	go func() {
		defer wg.Done()
		a = ParallelMerge(items[:middle])
	}()
	b := ParallelMerge(items[middle:])

	wg.Wait()
	return merge(a, b)
}
