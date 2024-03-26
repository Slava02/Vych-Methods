package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

var filepath = "C:\\Users\\Слава\\Desktop\\Виртуалка\\Vych Methods\\lab1\\test"

type matrix [][]float64
type vector []float64

type linearSystem struct {
	a matrix
	b vector
}

func main() {
	ls2 := parseData()
	res := SolveSystem(ls2.a, ls2.b)
	fmt.Println(res)
}

func parseData() linearSystem {
	f, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = f.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	s := bufio.NewScanner(f)
	s.Split(bufio.ScanWords)
	var size int
	fmt.Fscan(f, &size)
	ls := linearSystem{a: make(matrix, size), b: make(vector, size)}

	for i := 0; i < size; i++ {
		ls.a[i] = make(vector, size)
		for j := 0; j < size; j++ {
			s.Scan()
			ls.a[i][j], _ = strconv.ParseFloat(s.Text(), 64)
		}
		s.Scan()
		ls.b[i], _ = strconv.ParseFloat(s.Text(), 64)
	}

	err = s.Err()
	if err != nil {
		log.Fatal(err)
	}
	return ls
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
		fmt.Printf("forward elimanation (%d): \n", i)
		printMatrix(a)
		bufio.NewReader(os.Stdin).ReadBytes('\n')
	}
	return a
}

// backSubstitution makes backward substitution of jordan-gauss elimination
func backSubstitution(a matrix, m int) vector {
	x := make(vector, m)
	//fmt.Println(x)
	for i := m - 1; i >= 0; i-- {
		x[i] = a[i][m]
		//fmt.Printf("i=%d | x[i] = %.2f ", i, a[i][m])
		//fmt.Println(x)
		for j := i + 1; j < m; j++ {
			//ti, tj := x[i], x[j]
			x[i] -= a[i][j] * x[j]
			//fmt.Printf("i=%d , j=%d | %.2f -= %.2f * %.2f | x[i] = %.2f ", i, j, ti, a[i][j], tj, x[i])
			//fmt.Println(x)
		}
		//ti := x[i]
		x[i] /= a[i][i]
		//fmt.Printf("i=%d | %.2f /= %.2f | x[i]=%.2f ", i, ti, a[i][i], x[i])
		//fmt.Println(x)
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
