package wrap_test

import (
	"testing"

	"github.com/bbrks/wrap/v2"
)

func benchmarkWrap(b *testing.B, limit int) {
	w := wrap.NewWrapper()

	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		w.Wrap(loremIpsums[0], limit)
	}
}

func BenchmarkWrap0(b *testing.B)    { benchmarkWrap(b, 0) }
func BenchmarkWrap1(b *testing.B)    { benchmarkWrap(b, 1) }
func BenchmarkWrap2(b *testing.B)    { benchmarkWrap(b, 2) }
func BenchmarkWrap5(b *testing.B)    { benchmarkWrap(b, 5) }
func BenchmarkWrap10(b *testing.B)   { benchmarkWrap(b, 10) }
func BenchmarkWrap25(b *testing.B)   { benchmarkWrap(b, 25) }
func BenchmarkWrap80(b *testing.B)   { benchmarkWrap(b, 80) }
func BenchmarkWrap120(b *testing.B)  { benchmarkWrap(b, 120) }
func BenchmarkWrap500(b *testing.B)  { benchmarkWrap(b, 500) }
func BenchmarkWrap1000(b *testing.B) { benchmarkWrap(b, 1000) }
func BenchmarkWrap5000(b *testing.B) { benchmarkWrap(b, 5000) }

func benchmarkWrapOptimal(b *testing.B, limit int) {
	w := wrap.NewWrapper()
	w.MinimumRaggedness = true

	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		w.Wrap(loremIpsums[0], limit)
	}
}

func BenchmarkWrapOptimal10(b *testing.B)   { benchmarkWrapOptimal(b, 10) }
func BenchmarkWrapOptimal25(b *testing.B)   { benchmarkWrapOptimal(b, 25) }
func BenchmarkWrapOptimal80(b *testing.B)   { benchmarkWrapOptimal(b, 80) }
func BenchmarkWrapOptimal120(b *testing.B)  { benchmarkWrapOptimal(b, 120) }
func BenchmarkWrapOptimal500(b *testing.B)  { benchmarkWrapOptimal(b, 500) }
func BenchmarkWrapOptimal1000(b *testing.B) { benchmarkWrapOptimal(b, 1000) }

// BenchmarkGreedyVsOptimal compares greedy and optimal algorithms side by side.
func BenchmarkGreedyVsOptimal(b *testing.B) {
	input := loremIpsums[0]
	limit := 80

	b.Run("Greedy", func(b *testing.B) {
		w := wrap.NewWrapper()
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			w.Wrap(input, limit)
		}
	})

	b.Run("Optimal", func(b *testing.B) {
		w := wrap.NewWrapper()
		w.MinimumRaggedness = true
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			w.Wrap(input, limit)
		}
	})
}
