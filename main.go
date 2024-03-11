package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Введите квадратное уравнение формата: a*x^2 + b*x + c = 0")
	var a, b, c float64
	fmt.Scan(&a, &b, &c)
	fmt.Printf("Вы ввели %.2f*x^2 + %.2f*x + %.2f = 0", a, b, c)
	D := math.Pow(b, 2) - 4*a*c
	if D < 0 {
		fmt.Print("\nКорней нет")
	} else if D > 0 {
		fmt.Printf("\nx1 = %.2f\nx2 = %.2f", (-b-math.Sqrt(D))/(2*a), (-b+math.Sqrt(D))/(2*a))
	} else {
		fmt.Printf("\nx = %f.2", -b/(2*a))
	}
}
