package sliceh

import "math/rand"

// Chunk will separate an array by the given number of size
func Chunk[T any](items []T, size int) (chunks [][]T) {
	for size < len(items) {
		items, chunks = items[size:], append(chunks, items[0:size:size])
	}
	return append(chunks, items)
}

// Shuffle shuffles an array or slice and returns the shuffled slice or array
func Shuffle[T any](src []T) []T {
	dest := make([]T, len(src))
	perm := rand.Perm(len(src))
	for i, v := range perm {
		dest[v] = src[i]
	}
	return dest
}

// Filter returns only items that satisfy the given predicate (return the true predicate condition only)
func Filter[T any](items []T, callback func(T) bool) []T {
	matchItems := []T{}
	for _, item := range items {
		if callback(item) {
			matchItems = append(matchItems, item)
		}
	}
	return matchItems
}

// Map mapping slice of T
func Map[T1, T2 any](items []T1, callback func(T1) T2) []T2 {
	var resultItems []T2
	for _, item := range items {
		resultItems = append(resultItems, callback(item))
	}
	return resultItems
}

// Contains check whether the given items contains the given predicate
func Contains[T comparable](items []T, predicate T) bool {
	for _, item := range items {
		if item == predicate {
			return true
		}
	}
	return false
}

// NotContains check whether the given items not contains the given predicate
func NotContains[T comparable](items []T, predicate T) bool {
	for _, item := range items {
		if item == predicate {
			return false
		}
	}
	return true
}

// Random return a random T from the given slice of T
func Random[T any](slice ...T) T {
	if len(slice) > 1 {
		return slice[rand.Intn(len(slice)-1)]
	} else {
		return slice[0]
	}
}

// Merge merges two slices
func Merge[T any](first, second []T) []T {
	return append(first, second...)
}

// FirstOrNil returns the first item from a slice or nil if no such item exists
func FirstOrNil[T any](items []T) *T {
	length := len(items)
	if length == 0 {
		return nil
	}
	return &items[0]
}

// Last returns the last slice element
func Last[T any](items []T) T {
	length := len(items)
	return items[length-1]
}
