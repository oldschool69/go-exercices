package main

import "testing"

type unitTest struct {
	input []float64
	result float64
}

func TestAverage(t *testing.T) {

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