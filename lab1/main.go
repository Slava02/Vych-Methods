package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strconv"
)

var filepath = "C:\\Users\\Слава\\Desktop\\Виртуалка\\Vych Methods\\lab1\\test"

type Matrix [][]float64
type Vector []float64

type linearSystem struct {
	a Matrix
	b Vector
}

func main() {
	f := openFile(filepath)
	ls2 := parseData(f)
	res := SolveSystem(ls2.a, ls2.b)
	fmt.Println(res)
}

func openFile(name string) *os.File {
	f, err := os.Open(name)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = f.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	return f
}

func parseData(f io.Reader) linearSystem {
	s := bufio.NewScanner(f)
	defer func() {
		if err := s.Err(); err != nil {
			log.Fatal(err)
		}
	}()
	s.Split(bufio.ScanWords)
	var size int
	fmt.Fscan(f, &size)
	ls := linearSystem{a: make(Matrix, size), b: make(Vector, size)}
	for i := 0; i < size; i++ {
		ls.a[i] = make(Vector, size)
		for j := 0; j < size; j++ {
			s.Scan()
			ls.a[i][j], _ = strconv.ParseFloat(s.Text(), 64)
		}
		s.Scan()
		ls.b[i], _ = strconv.ParseFloat(s.Text(), 64)
	}
	return ls
}

// augmentedmatrix creates an augmented Matrix from the given Matrix and Vector.
func augmentedmatrix(a0 Matrix, b0 Vector) Matrix {
	m := len(b0)
	a := make(Matrix, m)

	for i, ai := range a0 {
		row := make(Vector, m+1)
		copy(row, ai)
		row[m] = b0[i]
		a[i] = row
	}
	return a
}

// SolveSystem solves given linear equations with Jordan-Gauss method
func SolveSystem(a0 Matrix, b0 Vector) Vector {
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
func replaceWithMaxRow(m Matrix, col int) {
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
func gaussianElimination(a Matrix, k, m int) Matrix {
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
func backSubstitution(a Matrix, m int) Vector {
	x := make(Vector, m)
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

// printMatrix prints given Matrix
func printMatrix(m Matrix) {
	for i, _ := range m {
		fmt.Printf("(%d) |", i)
		for _, col := range m[i] {
			fmt.Printf("%.2f ", col)
		}
		fmt.Println("|")
	}
	fmt.Println()
}
