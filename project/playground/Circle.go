package main

import "math"

type Circle struct {
	radius float64
}

func(c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func(c Circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}