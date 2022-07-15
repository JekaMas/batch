```go
package pipeline

import (
	"fmt"

	"github.com/JekaMas/batch"
)

type test struct {
	n int
	b bool
}

func TestBatch() {
	b := batch.NewBatch([]*test{
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

	fmt.Println("an incorrect sum", sum)
}

```