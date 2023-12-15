package main

import (
	"fmt"
	"math"
)

func f(x float64) float64 {
	return 2*math.Pow(x, 2) + 3*x - 5
}
func division(a, b, e float64) float64 {
	for math.Abs(b-a) > e {
		c := (a + b) / 2
		if f(a)*f(c) >= 0 {
			a = c
		} else {
			b = c
		}
	}
	return (a + b) / 2
}

func secant(a, b, e float64) float64 {
	xn := 0.
	for {
		y := xn
		xn = b - ((b-a)/(f(b)-f(a)))*f(b)
		a = b
		b = xn
		if math.Abs(y-xn) < e {
			break
		}
	}
	return xn
}

// func secant(a, b, e float64) float64 {
// 	for math.Abs(b-a) > e {
// 		a -= (b - a) * f(a) / (f(b) - f(a))
// 		b -= (a - b) * f(b) / (f(a) - f(b))
// 	}
// 	return b
// }

func main() {
	var (
		a   float64
		b   float64
		e   float64
		opt int
	)
	fmt.Print("Левая граница:")
	fmt.Scan(&a)
	fmt.Print("Правая граница:")
	fmt.Scan(&b)
	fmt.Print("Точность:")
	fmt.Scan(&e)
	fmt.Println("Выберете метод:")
	fmt.Println("1. Метод деления пополам:")
	fmt.Println("2. Метод секущих:")
	fmt.Scan(&opt)
	if opt == 1 {
		fmt.Println(division(a, b, e))
	} else {
		fmt.Println(secant(a, b, e))
	}
}
