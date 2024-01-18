package main

import "fmt"

type person struct {
	name string
	age int
}

type dict struct {
	data map[int]string
}


func newDict() *dict {
	d := dict{}
	d.data = map[int]string{}
	return &d
}

func main() {


	p := person{}
	p.name = "Kim"
	p.age = 20

	p1 := person{name: "Lee", age: 10}
	p2 := person{"Ahn", 22}

	fmt.Println(p)
	fmt.Println(p1)
	fmt.Println(p2)

	p3 := new(person)
	p3.name = "Son"  //포인터여도 .을 사용한다
	p3.age = 50

	fmt.Println(p3) 

	dic := newDict()
	dic.data[1] = "A"
}

// struct 개체를 다른 함수의 파라미터로 넘긴다면, Pass by Value에 따라 객체를 복사해서 전달하게 된다. 
// 그래서 만약 Pass by Reference로 struct를 전달하고자 한다면, struct의 포인터를 전달해야 한다.