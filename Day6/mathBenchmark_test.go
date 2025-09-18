//Use Benchmark at the start of function.
//run Benchmark testing with `go test -bench=.`

package mathutils

import "testing"

func BenchmarkAdd(b *testing.B) {
	for i := 1; i < b.N; i++ {
		Add(100, 200)
	}
}
