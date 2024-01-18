package main

import (
	"testing"
	"math"
)


func TestShape(t *testing.T) {
	r := Rect{10., 20.}
	c := Circle{10.}

	checkArea(t, r, c)
}

func checkArea(t *testing.T, shapes ...Shape) {
	s1 := shapes[0]
	s2 := shapes[1]

	if s1.area() != 200.0 {
		t.Errorf("Rect area() is wrong %f", s1.area())
	}

	if s1.perimeter() != 400.0 {
		t.Errorf("Rect perimeter() is wrong  %f", s1.perimeter())
	}

	if math.Round(s2.area()) != 314.0 {
		t.Errorf("Circle area() is wrong %f", s2.area())
	}

	delta := 0.001 // 허용 오차값
	if math.Abs(math.Round(s2.perimeter())-62.8) < delta {
    	t.Errorf("Circle perimeter() is wrong. Expected %v, got %v", 62.8, s2.perimeter())
	}
}