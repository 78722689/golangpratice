package gc

import "testing"

func Benchmark_1_GetFieldValue(b *testing.B) {

	/*var x int
	for i := 0; i < b.N; i++ {
		x++
	}*/
	GCMain(1)

}

func Benchmark_2_GetFieldValue(b *testing.B) {

	/*var x int
	for i := 0; i < b.N; i++ {
		x++
	}*/
	GCMain(10)

}
