package main

import "fmt"

//fmt 패키지를 사용하기 위해 import문을 사용한다.
//출력에 대한 부분은 c언어와 비슷하다.

//입력 부분을 중점적으로 다룬다

func Scan() {
	var a int = 0
	var b int = 0
	n, err := fmt.Scan(&a, &b)

	if err != nil {
		fmt.Println(n, err)

	} else {
		fmt.Println(n, a, b)

	}
}

func main() {
	//scan 함수를 매개변수로 포인터를 받는다.}
	Scan()
	mustInt()
	clearStream()
	//입력을 받아서 출력하는 코드이다.
}
