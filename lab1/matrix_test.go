package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestGauss(t *testing.T) {
	var tests = []struct {
		name string
		a    Matrix
		b    Vector
		want Vector
	}{
		{
			name: "default",
			a: Matrix{
				{4.0, -1.0, 1.0},
				{1.0, 6.0, 2.0},
				{-1.0, -2, 5.0}},
			b:    Vector{4.0, 9.0, 2.0},
			want: Vector{1.0, 1.0, 1.0},
		},
		{
			name: "zero",
			a: [][]float64{
				{0.0, 0.0, 0.0},
				{0.0, 0.0, 0.0},
				{0.0, 0, 0.0}},
			b:    []float64{0.0, 0.0, 0.0},
			want: []float64{0.0, 0.0, 0.0},
		},
	}
	for _, test := range tests {
		name := fmt.Sprintf("case %s:", test.name)
		t.Run(name, func(t *testing.T) {
			got := SolveSystem(test.a, test.b)
			if !equalVec(got, test.want) {
				t.Errorf("got %v, want %v", got, test.want)
			}
		})
	}
}

func TestParseData(t *testing.T) {
	var tests = []struct {
		name, read string
		want       linearSystem
	}{
		{
			name: `default`,
			read: `3
			4.0 -1.0 1.0 4.0
			1.0 6.0 2.0 9.0
			-1.0 -2 5.0 2.0`,
			want: linearSystem{
				a: Matrix{
					{4.0, -1.0, 1.0},
					{1.0, 6.0, 2.0},
					{-1.0, -2, 5.0}},
				b: Vector{4.0, 9.0, 2.0},
			},
		},
	}
	for _, test := range tests {
		name := fmt.Sprintf("case %s:", test.name)
		t.Run(name, func(t *testing.T) {
			got := parseData(strings.NewReader(test.read))
			if !equalLS(got, test.want) {
				t.Errorf("got %v, want %v", got, test.want)
			}
		})
	}
}

func equalVec(a Vector, b Vector) bool {
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func equalMatrix(a Matrix, b Matrix) bool {
	for i, v := range a {
		if !equalVec(v, b[i]) {
			return false
		}
	}
	return true
}

func equalLS(a linearSystem, b linearSystem) bool {
	return equalMatrix(a.a, b.a) && equalVec(a.b, b.b)
}
