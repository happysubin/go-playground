package main

type Rect struct {
	width int 
	height int
}

func (r Rect) area() int {
	return r.width * r.height
}