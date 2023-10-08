package main

import (
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

func Jakoby(a [][]float64, b []float64, eps float64) []float64 {
	tempX := make([]float64, len(a))

	// Начальное приближение
	x := make([]float64, len(a))
	for i := 0; i < len(a); i++ {
		x[i] = b[i] / a[i][i]
	}
	norma := 0.0
	for {
		for i := 0; i < len(a); i++ {
			tempX[i] = b[i]
			for j := 0; j < len(a); j++ {
				if i != j {
					tempX[i] -= a[i][j] * x[j]
				}
			}
			tempX[i] /= a[i][i]
		}

		// Норма
		norma = math.Abs(x[0] - tempX[0])
		for h := 0; h < len(a); h++ {
			if math.Abs(x[h]-tempX[h]) > norma {
				norma = math.Abs(x[h] - tempX[h])
			}
			x[h] = tempX[h]
		}
		// Условие окончания цикла
		if norma < eps {
			break
		}
	}
	return x
}

func Seidel(a [][]float64, b []float64, eps float64) []float64 {
	tempX := make([]float64, len(a))
	norma := 0.0
	x := make([]float64, len(a))
	for i := 0; i < len(a); i++ {
		x[i] = b[i] / a[i][i]
	}
	for {
		for i := 0; i < len(a); i++ {
			tempX[i] = b[i]
			for j := 0; j < len(a); j++ {
				if i != j {
					tempX[i] -= a[i][j] * tempX[j]
				}
			}
			tempX[i] /= a[i][i]
		}
		norma = math.Abs(x[0] - tempX[0])
		for h := 0; h < len(a); h++ {
			if math.Abs(x[h]-tempX[h]) > norma {
				norma = math.Abs(x[h] - tempX[h])
			}
			x[h] = tempX[h]
		}
		if norma < eps {
			break
		}
	}
	return x
}

func GetMatrix() ([][]float64, []float64) {
	file, err := os.Open("matrix.txt")
	if err != nil {
		fmt.Println(err)
	}
	text, err := io.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	rows := strings.Split(string(text), "\n")
	var matrixStr [][]string
	for _, row := range rows {
		matrixStr = append(matrixStr, strings.Split(row, " "))
	}
	matrix := make([][]float64, len(matrixStr))
	for i := 0; i < len(matrixStr); i++ {
		matrix[i] = make([]float64, len(matrixStr[i]))
		for j := 0; j < len(matrixStr[i]); j++ {
			matrix[i][j], _ = strconv.ParseFloat(matrixStr[i][j], 64)
		}
	}
	n := len(matrix)
	a := make([][]float64, n)
	b := make([]float64, n)
	for i := 0; i < n; i++ {
		a[i] = make([]float64, n)
		b[i] = matrix[i][n]
		for j := 0; j < n; j++ {
			a[i][j] = matrix[i][j]
		}
	}

	return a, b
}

func main() {
	a, b := GetMatrix()
	x := Jakoby(a, b, 0.001)
	for _, i := range x {
		fmt.Printf("%.4f ", i)
	}
}
