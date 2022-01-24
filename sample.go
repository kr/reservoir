// Package reservoir samples values uniformly at random,
// without replacement, from an unbounded sequence of inputs.
// It provides a representative sample
// when the sequence has unknown length
// or is too big to store in its entirety.
//
// The approach used here is commonly known as reservoir sampling.
package reservoir

import "math/rand"

// A Sample[T] collects a fixed number of items,
// chosen uniformly at random,
// from an unbounded input sequence.
// The number of items collected is its sampling capacity.
//
// The zero value is a valid Sample
// with a sampling capacity of 0.
// (It will not sample any items.)
type Sample[T any] struct {
	n   int
	data []T // slice header is constant
}

// New returns a new Sample
// with sampling capacity cap.
func New[T any](cap int) *Sample[T] {
	return &Sample[T]{data: make([]T, cap)}
}

// Reset empties s.
func (s *Sample[T]) Reset() {
	s.n = 0
}

// Add adds v to s with a probability
// adjusted so that the contents of s
// at any time are chosen uniformly at random
// from the inputs so far.
func (s *Sample[T]) Add(v T) {
	if s.n < len(s.data) {
		s.data[s.n] = v
	} else if i := rand.Intn(s.n+1); i < len(s.data) {
		// Sample v with probability len(s.data)/n
		// (where n is the number of items so far, including v).
		// Replace a sampled item with prob. 1/len(s.data).
		// See Jeffrey S. Vitter, Random sampling with a reservoir,
		// ACM Trans. Math. Softw. 11 (1985), no. 1, 37–57.
		s.data[i] = v
	}
	s.n++
}

// Read reads the current contents of s into p.
// It returns the number of values read,
// at most the sampling capacity of s.
func (s *Sample[T]) Read(p []T) int {
	d := s.data
	if s.n < len(d) {
		d = d[:s.n]
	}
	return copy(p, d)
}

// Cap returns the sampling capacity of s.
func (s *Sample[T]) Cap() int {
	return len(s.data)
}

// Added returns the number of calls to Add.
func (s *Sample[T]) Added() int {
	return s.n
}
