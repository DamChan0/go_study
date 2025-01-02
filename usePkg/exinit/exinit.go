package exinit

import (
	"fmt"
)

var (
	a = c + b
	b = f()
	c = f()
	d = 3
)

func init() { // init 함수는 패키지가 로드될 때 가장 먼저 실행되는 함수
	d++
	fmt.Println("extint.init function ", d)

}

func f() int {
	d++
	fmt.Println("f() function ", d)
	return d
}

func PrintD() {
	fmt.Println("d : ", d)
}
