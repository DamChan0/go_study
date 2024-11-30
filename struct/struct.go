package main

import (
	"fmt"
	a "go_study/Struct/A"
	em "go_study/Struct/em"
)

// 선언 및 기본 사용

type Person struct {
	name   string
	age    int
	skills []string
	career int
}

func main() {
	println("Struct Example")
	// 구조체 선언
	var p1 Person // 기본값으로 초기본

	fmt.Println(p1) // { 0  0}
	p1.name = "Gopher"
	p1.age = 10
	p1.skills = []string{"Golang"}
	p1.career = 1

	// 아래 두 출력은 동일한 결과를 출력한다.
	fmt.Printf("%v\n", p1)
	fmt.Println(p1)

	var p2 Person = Person{
		name:   "DamChan",
		age:    28,
		skills: []string{"C++", "c", "Golang", "python"},
		career: 2,
	}

	fmt.Println(p2)

	var p3 Person = Person{
		"DamChan",
		28,
		[]string{"C++", "c", "Golang", "python"},
		2,
	}
	fmt.Println(p3)
	a.A()
	em.EmbeddedStruct()
}
