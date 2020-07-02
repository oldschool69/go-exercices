package carecao

import (
	"fmt"
	"testing"
)


func TestAverage(t *testing.T) {

	type unitTest struct {
		input []float64
		result float64
	}

	tcs := []unitTest{
		{
			[]float64{1, 2},
			1.5,
		},
		{
			[]float64{2, 2, 2},
			2,
		},
		{
			[]float64{1, 2, 3, 4, 5},
			3,
		},
	}

	for _, v := range tcs {
		avg := Average(v.input...)

		if avg != v.result {
			t.Error("Expected", v.result, "Got", avg)
		}
	}

}

func ExampleAverage() {
	values := []float64{2,3,10}
	fmt.Println(Average(values...))
	// Output:
	// 5
}

func BenchmarkAverage(b *testing.B) {
	values := []float64{2,3,10}
	for i := 0; i < b.N; i++ {
		Average(values...)
	}
}