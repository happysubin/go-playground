package main

import "testing"


func TestRect(t *testing.T) { //TestXxx와 같은 이름으로 작성해야한다.
	rect := Rect{10, 20}

	area := rect.area()

	want := 200.0
	if (area != want) {
		t.Errorf("result is wrong")
	}
}


