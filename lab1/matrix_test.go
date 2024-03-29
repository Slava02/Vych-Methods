package main

import (
	"fmt"
	"testing"
)

func TestGauss(t *testing.T) {
	var tests = []struct {
		name string
		a    matrix
		b    vector
		want vector
	}{
		{
			name: "default",
			a: [][]float64{
				{4.0, -1.0, 1.0},
				{1.0, 6.0, 2.0},
				{-1.0, -2, 5.0}},
			b:    []float64{4.0, 9.0, 2.0},
			want: []float64{1.0, 1.0, 1.0},
		},
	}
	for _, test := range tests {
		name := fmt.Sprintf("case %s:", test.name)
		t.Run(name, func(t *testing.T) {
			got := SolveSystem(test.a, test.b)
			if !equal(got, test.want) {
				t.Errorf("got %v, want %v", got, test.want)
			}
		})
	}
}

func equal(a vector, b vector) bool {
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
