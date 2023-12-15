package main

import (
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

func Gauss(Matrix [][]float64) [][]float64 {
	n := len(Matrix)

	xirtaM := make([][]float64, n)
	for i := range xirtaM {
		xirtaM[i] = make([]float64, n)
		xirtaM[i][i] = 1
	}

	Matrix_Big := make([][]float64, n)
	for i := range Matrix_Big {
		Matrix_Big[i] = make([]float64, 2*n)
		for j := range Matrix_Big[i] {
			if j < n {
				Matrix_Big[i][j] = Matrix[i][j]
			} else {
				Matrix_Big[i][j] = xirtaM[i][j-n]
			}
		}
	}

	for k := 0; k < n; k++ {
		for i := 0; i < 2*n; i++ {
			Matrix_Big[k][i] = Matrix_Big[k][i] / Matrix[k][k]
		}
		for i := k + 1; i < n; i++ {
			K := Matrix_Big[i][k] / Matrix_Big[k][k]
			for j := 0; j < 2*n; j++ {
				Matrix_Big[i][j] = Matrix_Big[i][j] - Matrix_Big[k][j]*K
			}
		}
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				Matrix[i][j] = Matrix_Big[i][j]
			}
		}
	}

	for k := n - 1; k > -1; k-- {
		for i := 2*n - 1; i > -1; i-- {
			Matrix_Big[k][i] = Matrix_Big[k][i] / Matrix[k][k]
		}
		for i := k - 1; i > -1; i-- {
			K := Matrix_Big[i][k] / Matrix_Big[k][k]
			for j := 2*n - 1; j > -1; j-- {
				Matrix_Big[i][j] = Matrix_Big[i][j] - Matrix_Big[k][j]*K
			}
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			xirtaM[i][j] = Matrix_Big[i][j+n]
		}
	}

	return xirtaM
}

func GetMatrix() [][]float64 {
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

	return matrix
}

func main() {
	a := GetMatrix()
	b := Gauss(a)
	f, _ := os.Create("result.txt")
	for i := 0; i < len(b); i++ {
		for j := 0; j < len(b[i]); j++ {
			if math.IsNaN(b[i][j]) {
				f, _ := os.Create("result.txt")
				f.Write([]byte("Ошибка"))
				return
			}
			f.Write([]byte(fmt.Sprintf("%.2f\t", b[i][j])))
		}
		f.Write([]byte("\n"))
	}
}
