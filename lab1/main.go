package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type matrix [][]float64
type vector []float64

type linearSystem struct {
	a matrix
	b vector
}

var ls = linearSystem{
	a: [][]float64{
		{3.0, 2.0, -4.0},
		{2.0, 3.0, 3.0},
		{5.0, -3, 1.0},
	},
	b: []float64{3.0, 15.0, 14.0},
}

// SOLUTION: 3 1 2

func main() {
	res := SolveSystem(ls.a, ls.b)
	fmt.Println(res)
}

// augmentedmatrix creates an augmented matrix from the given matrix and vector.
func augmentedmatrix(a0 matrix, b0 vector) matrix {
	m := len(b0)
	a := make(matrix, m)

	for i, ai := range a0 {
		row := make(vector, m+1)
		copy(row, ai)
		row[m] = b0[i]
		a[i] = row
	}
	return a
}

// SolveSystem solves given linear equations with Jordan-Gauss method
func SolveSystem(a0 matrix, b0 vector) vector {
	a := augmentedmatrix(a0, b0)
	printMatrix(a)
	for i, _ := range a {
		replaceWithMaxRow(a, i)
		//printMatrix(a)
		a = gaussianElimination(a, i, len(a))
	}
	return backSubstitution(a, len(a))
}

// replaceWithMaxRow replaces m[col] row with the one of the underlying rows with the modulo greatest first element.
func replaceWithMaxRow(m matrix, col int) {
	maxElem := m[col][col]
	maxRow := col
	for i := col + 1; i < len(m[0])-1; i++ {
		if math.Abs(m[i][col]) > math.Abs(maxElem) {
			maxElem = m[i][col]
			maxRow = i
		}
	}
	if maxRow != col {
		m[col], m[maxRow] = m[maxRow], m[col]
		fmt.Printf("Max elem:%f | raw(%d) <--> raw(%d)\n", maxElem, col, maxRow)
		printMatrix(m)
	} else {
		fmt.Printf("Max elem:%f\n", maxElem)
	}
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}

// gaussianElimination makes iteration of jordan-gauss elimination
func gaussianElimination(a matrix, k, m int) matrix {
	for i := k + 1; i < m; i++ {
		for j := k + 1; j <= m; j++ {
			a[i][j] -= a[k][j] * (a[i][k] / a[k][k])
		}
		a[i][k] = 0
		fmt.Printf("elimanation (%d): \n", i)
		printMatrix(a)
		bufio.NewReader(os.Stdin).ReadBytes('\n')
	}
	return a
}

// backSubstitution makes backward substitution of jordan-gauss elimination
func backSubstitution(a matrix, m int) vector {
	x := make(vector, m)
	for i := m - 1; i >= 0; i-- {
		x[i] = a[i][m]
		for j := i + 1; j < m; j++ {
			x[i] -= a[i][j] * x[j]
		}
		x[i] /= a[i][i]
	}
	return x
}

func printMatrix(m matrix) {
	for i, _ := range m {
		fmt.Printf("(%d) |", i)
		for _, col := range m[i] {
			fmt.Printf("%.2f ", col)
		}
		fmt.Println("|")
	}
	fmt.Println()
}
