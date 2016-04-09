package tempconv

import (
	"math"
	"testing"
)

func TestCToK(t *testing.T) {
	var tests = []struct {
		input Celsius
		want  Kelvin
	}{
		{0, 273.15},
		{100, 373.15},
		{-100, 173.15},
		{-273.15, 0},
	}
	for _, test := range tests {
		if got := CToK(test.input); math.Abs(float64(test.want-got)) > 0.00001 {
			t.Errorf("CToF(%v) = %v, wanted %v", test.input, got, test.want)
		}
	}
}
