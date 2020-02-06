package dynamic_programming

import "testing"

func BenchmarkGetClimbingWays(b *testing.B) {

	number := 20

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		GetClimbingWays3(number)
	}
}

func BenchmarkGetClimbingWays1(b *testing.B) {

	number := 20

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		GetClimbingWays1(number)
	}
}
