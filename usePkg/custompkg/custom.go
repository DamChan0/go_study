package custompkg

import (
	"fmt"
)

type Student struct {
	Name string
	Age  int
}

var PublicInt int = 10  // Public variable
var privateInt int = 20 // Private variable custompkg 를 import 한 곳에서 사용할 수 없음

func CustomFunc() {
	fmt.Println("This is a custom package")
}

func CustomFunc2() {
	fmt.Println("This is another custom package")
}
