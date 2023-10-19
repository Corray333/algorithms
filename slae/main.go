package main

import (
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

func Jakoby(a [][]float64, b []float64) []float64 {
	eps := 0.001
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

func Down(a [][]float64, nb []float64) []float64 {
	b := make([][]float64, len(nb))
	u := make([][]float64, len(nb))
	for i := 0; i < len(nb); i++ {
		b[i] = make([]float64, 1)
		b[i][0] = nb[i]
		u[i] = make([]float64, 1)
		u[i][0] = b[i][0] / a[i][i]
	}
	for k := 0; k < 100; k++ {
		r := multiplyMatrices(a, u)
		for i := 0; i < len(b); i++ {
			r[i][0] -= b[i][0]
		}
		delta := make([][]float64, len(b))
		for i := 0; i < len(b); i++ {
			delta[i] = make([]float64, 1)
			delta[i][0] = 2 * r[i][0]
		}
		ar := multiplyMatrices(a, r)
		alpha := multiplyVectors(ar, r) / (2 * multiplyVectors(ar, ar))
		for i := 0; i < len(b); i++ {
			u[i][0] -= delta[i][0] * alpha
		}
	}
	result := make([]float64, len(u))
	for i := 0; i < len(b); i++ {
		result[i] = u[i][0]
	}
	return result
}

func multiplyMatrices(matrix1 [][]float64, matrix2 [][]float64) [][]float64 {
	rows1 := len(matrix1)
	cols1 := len(matrix1[0])
	rows2 := len(matrix2)
	cols2 := len(matrix2[0])

	// Проверяем, можно ли умножить матрицы
	if cols1 != rows2 {
		fmt.Println("Невозможно умножить матрицы")
		return nil
	}

	// Создаем результирующую матрицу
	result := make([][]float64, rows1)
	for i := range result {
		result[i] = make([]float64, cols2)
	}

	// Умножаем матрицы
	for i := 0; i < rows1; i++ {
		for j := 0; j < cols2; j++ {
			for k := 0; k < cols1; k++ {
				result[i][j] += matrix1[i][k] * matrix2[k][j]
			}
		}
	}

	return result
}

func multiplyVectors(v1 [][]float64, v2 [][]float64) float64 {
	res := 0.
	for i := 0; i < len(v1); i++ {
		res += v1[i][0] * v2[i][0]
	}
	return res
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
	x := Jakoby(a, b)
	file, _ := os.Create("result.txt")
	for _, i := range x {
		file.Write([]byte(fmt.Sprintf("%.4f ", i)))
	}
	file.Close()
}
