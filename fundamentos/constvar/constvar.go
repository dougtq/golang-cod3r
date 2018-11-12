package main

import (
	"fmt"
	m "math"
)

func main() {
	const Pi float64 = 3.1415
	var raio = 3.2

	// forma reduzida de criar uma var
	area := Pi * m.Pow(raio, 2)

	fmt.Println("A área da circunferência é", area)

	const (
		a = 1
		b = 2
	)

	var (
		d = 'a'
		c = 'd'
	)

	fmt.Println(b, c, a, d)
}
