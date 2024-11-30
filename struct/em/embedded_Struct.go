package em

import "fmt"

// 선언 및 기본 사용
func EmbeddedStruct() {

	type Address struct {
		load string
		city string
		zip  int
	}

	// 실제 타입 지정은 맨 마지막에 한다.
	type Person struct {
		name    string
		age     int
		address Address
		skills  []string
		career  int
	}

	var p1 Person
	p1.name = "Gopher"
	p1.age = 10
	p1.skills = []string{"Golang"}
	p1.career = 1
	p1.address = Address{
		load: "Seoul",
		city: "Gangnam",
		zip:  12345,
	}

	fmt.Println(p1)

}
