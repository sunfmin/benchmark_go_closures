package benchmark_go_closures

import (
	"testing"
	// "fmt"
)

func BenchmarkClosureWithPassedIn(b *testing.B) {
	for i := 0; i < b.N; i++ {
		f := func(i int) (r int) {
			// fmt.Sprintf("Hello %d", i)
			r = i * 2
			return
		}
		f(i)
	}
}

func BenchmarkClosureWithGlobalVariable(b *testing.B) {
	for i := 0; i < b.N; i++ {
		f := func() (r int) {
			// fmt.Sprintf("Hello %d", i)
			r = i * 2
			return
		}
		f()
	}
}

func BenchmarkNormalFunctionCall(b *testing.B) {
	for i := 0; i < b.N; i++ {
		multiple2(i)
	}
}

func multiple2(i int) (r int) {
	// fmt.Sprintf("Hello %d", i)
	r = i * 2
	return
}
