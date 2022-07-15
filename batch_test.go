package batch

import (
	"fmt"
	"testing"
)

type test struct {
	n int
	b bool
}

func TestBatch(t *testing.T) {
	t.Parallel()

	b := NewBatch([]*test{
		{50, false},
		{3, true},
	}...)

	var sum int

	b.Apply(
		b.NewWrite(func(vs ...*test) {
			p := vs[0]
			q := vs[1]

			if p != nil && q != nil && *p != *q {
				*p = *q
			}
		}),
		b.NewRead(func(vs ...*test) {
			sum = vs[0].n + vs[1].n
		}),
	)

	if sum != 53 && sum != 6 {
		t.Fatal(t, fmt.Sprintf("wasn't serialized: got %d, expected 53 or 6", sum))
	}
}
