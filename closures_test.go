package benchmark_go_closures

import (
	"fmt"
	"testing"
)

func BenchmarkAnonymousFunction(b *testing.B) {
	var t int64
	for i := 0; i < b.N; i++ {
		f := func(j int64) (r int64) {
			r = j * 2
			return
		}

		t = f(int64(i))
	}
	fmt.Sprintln(t)
}

func BenchmarkReusingAnonymousFunction(b *testing.B) {
	var t int64
	f := func(j int64) (r int64) {
		r = j * 2
		return
	}

	for i := 0; i < b.N; i++ {
		t = f(int64(i))
	}

	fmt.Sprintln(t)
}

func BenchmarkClosure(b *testing.B) {
	var t int64
	for i := 0; i < b.N; i++ {
		f := func() (r int64) {
			r = int64(i) * 2
			return
		}
		t = f()
	}
	fmt.Sprintln(t)

}

func BenchmarkReusingClosure(b *testing.B) {
	var t int64
	var j int64
	f := func() (r int64) {
		r = j * 2
		return
	}

	for i := 0; i < b.N; i++ {
		t = f()
	}
	fmt.Sprintln(t)

}

func BenchmarkNormalFunction(b *testing.B) {
	var t int64
	for i := 0; i < b.N; i++ {
		t = multiple2(int64(i))
	}
	fmt.Sprintln(t)
}

func multiple2(i int64) (r int64) {
	r = i * 2
	return
}
