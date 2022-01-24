package reservoir_test

import (
	"errors"
	"fmt"
	"math/rand"

	"kr.dev/reservoir"
)

func Example_errors() {
	// Make & populate a new reservoir of errors with capacity 5.
	rs := reservoir.New[error](5)
	rs.Add(errors.New("first"))
	rs.Add(errors.New("second"))
	rs.Add(errors.New("third"))

	// Read a sample into errs.
	errs := make([]error, rs.Cap())
	n := rs.Read(errs)

	fmt.Println("our sample:", errs[:n])
	// Output:
	// our sample: [first second third]
}

func Example_many() {
	rand.Seed(0)
	// Make & populate a new reservoir of ints with capacity 3.
	rs := reservoir.New[int](3)
	for i := 0; i < 10_000; i++ {
		rs.Add(i)
	}

	// Read a sample into ints.
	ints := make([]int, rs.Cap())
	n := rs.Read(ints)

	fmt.Println("our sample:", ints[:n])
	// Output:
	// our sample: [9539 5555 7160]
}
