package main

import (
	"errors"
	"fmt"
	"log"
	"math"
	"math/rand"
)

// // func Jacobi(a [][]float64, n int) ([][]float64, [][]float64) {
// // 	const maxIterations = 100
// // 	const eps = 1e-9

// // 	v := make([][]float64, n)
// // 	for i := range v {
// // 		v[i] = make([]float64, n)
// // 		for j := range v[i] {
// // 			if i == j {
// // 				v[i][j] = 1.0
// // 			}
// // 		}
// // 	}

// // 	for k := 0; k < maxIterations; k++ {
// // 		m := 0
// // 		n := 1
// // 		for i := 0; i < len(a); i++ {
// // 			for j := i + 1; j < len(a); j++ {
// // 				if math.Abs(a[i][j]) > math.Abs(a[m][n]) {
// // 					m = i
// // 					n = j
// // 				}
// // 			}
// // 		}

// // 		if math.Abs(a[m][n]) < eps {
// // 			return a, v
// // 		}

// // 		d := (a[n][n] - a[m][m]) / (2 * a[m][n])
// // 		t := 1.0 / (math.Abs(d) + math.Sqrt(d*d+1))
// // 		if d < 0 {
// // 			t = -t
// // 		}
// // 		c := 1.0 / math.Sqrt(t*t+1)
// // 		s := t * c
// // 		tau := s / (1.0 + c)
// // 		temp := a[m][n]
// // 		a[m][n] = 0
// // 		a[m][m] -= t * temp
// // 		a[n][n] += t * temp
// // 		for i := 0; i < m; i++ {
// // 			temp = a[i][m]
// // 			a[i][m] -= s * (a[i][n] + tau*temp)
// // 			a[i][n] += s * (temp - tau*a[i][m])
// // 		}
// // 		for i := m + 1; i < n; i++ {
// // 			temp = a[m][i]
// // 			a[m][i] -= s * (a[i][n] + tau*temp)
// // 			a[i][n] += s * (temp - tau*a[m][i])
// // 		}
// // 		for i := n + 1; i < len(a); i++ {
// // 			temp = a[m][i]
// // 			a[m][i] -= s * (a[n][i] + tau*temp)
// // 			a[n][i] += s * (temp - tau*a[m][i])
// // 		}
// // 		for i := 0; i < len(a); i++ {
// // 			temp = v[i][m]
// // 			v[i][m] -= s * (v[i][n] + tau*temp)
// // 			v[i][n] += s * (temp - tau*v[i][m])
// // 		}
// // 	}

// // 	fmt.Println("Jacobi method did not converge")
// // 	return nil, nil
// // }

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

func Rotations(matrix [][]float64, eps float64) ([][]float64, error) {
	if len(matrix) != len(matrix[0]) {
		return nil, errors.New("Метод вращений не подходит для данного типа матриц.")
	}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix); j++ {
			if matrix[i][j] != matrix[j][i] {
				return nil, errors.New("Метод вращений не подходит для данного типа матриц.")
			}
		}
	}
	for {
		max := matrix[0][1]
		i := 0
		j := 1
		for k := 0; k < len(matrix); k++ {
			for g := 0; g < len(matrix); g++ {
				if k != g {
					if math.Abs(matrix[k][g]) > max {
						max = matrix[k][g]
						i = k
						j = g
					}
				}
			}
		}
		if math.Abs(max) < eps {
			return matrix, nil
		}
		var tau float64
		if matrix[i][i] == matrix[j][j] {
			tau = math.Pi / 4
		} else {
			tau = math.Tan(2*matrix[i][j]/(matrix[i][i]-matrix[j][j])) / 2
		}
		R := make([][]float64, len(matrix))
		for k := 0; k < len(matrix); k++ {
			R[k] = make([]float64, len(matrix))
			R[k][k] = 1
		}
		R[i][i] = math.Cos(tau)
		R[j][j] = math.Cos(tau)
		R[i][j] = -math.Sin(tau)
		R[j][i] = math.Sin(tau)
		invR := inverseMatrix(R)
		matrix = multiplyMatrices(multiplyMatrices(invR, matrix), R)
	}
}

func det(matrix [][]float64) float64 {
	if len(matrix) == 2 {
		return matrix[0][0]*matrix[1][1] - matrix[0][1]*matrix[1][0]
	} else {
		sum := 0.
		for i := 0; i < len(matrix); i++ {
			minor := make([][]float64, len(matrix))
			for j := 0; j < len(matrix); j++ {
				minor[j] = make([]float64, 0)
				minor[j] = append(minor[j], matrix[j][1:]...)
			}
			minor = append(minor[:i], minor[i+1:]...)
			sign := 1.
			if i%2 == 1 {
				sign = -1
			}
			sum += sign * matrix[i][0] * det(minor)
		}
		return sum
	}
}

func transposeMatrix(matrix [][]float64) [][]float64 {
	transp := make([][]float64, len(matrix[0]))
	for i := 0; i < len(matrix[0]); i++ {
		transp[i] = make([]float64, len(matrix))
		for j := 0; j < len(matrix); j++ {
			transp[i][j] = matrix[j][i]
		}
	}
	return transp
}

func inverseMatrix(matrix [][]float64) [][]float64 {
	determinant := det(matrix)
	transposed := transposeMatrix(matrix)
	result := make([][]float64, len(transposed))
	for i := 0; i < len(transposed); i++ {
		result[i] = make([]float64, len(transposed[0]))
		for j := 0; j < len(transposed); j++ {
			minor := make([][]float64, len(transposed))
			for k := 0; k < len(transposed); k++ {
				minor[k] = make([]float64, 0)
				minor[k] = append(append(minor[k], transposed[k][:j]...), transposed[k][j+1:]...)
			}
			minor = append(minor[:i], minor[i+1:]...)
			sign := 1.
			if (i+j)%2 == 1 {
				sign = -1
			}
			result[i][j] = sign * det(minor) / determinant
		}
	}
	return result
}

func itaretions(matrix [][]float64) float64 {
	r := make([][]float64, len(matrix))
	for i := 0; i < len(matrix); i++ {
		r[i] = make([]float64, 1)
		r[i][0] = float64(rand.Int31())
	}
	mu := 0.
	for i := 0; i < 100; i++ {
		tempR := make([][]float64, 0)
		tempR = append(tempR, r...)
		r = multiplyMatrices(matrix, r)
		mu = multiplyVectors(tempR, r) / multiplyVectors(tempR, tempR)
	}
	return mu

}

func multiplyVectors(v1 [][]float64, v2 [][]float64) float64 {
	res := 0.
	for i := 0; i < len(v1); i++ {
		res += v1[i][0] * v2[i][0]
	}
	return res
}

func main() {
	m := [][]float64{
		{1, math.Sqrt(2), 2},
		{5, 3, math.Sqrt(2)},
		{2, math.Sqrt(2), 1},
	}
	// m := [][]float64{
	// 	{-1, -6},
	// 	{2, 6},
	// }
	res, err := Rotations(m, 0.001)
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m); j++ {
			fmt.Printf("%.3f\t", res[i][j])
		}
		fmt.Println()
	}
	// fmt.Print(itaretions(m))
}
