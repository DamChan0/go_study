package main

import "fmt"

func Add(a int, b int) int { // 매개변수 와 반환값을 정의한다와
	return a + b
}

func MultiReturn(a, b int) (int, bool) { // 매개변수와 반환값을 정의한다.
	if b == 0 {
		return 0, false
	}
	return a / b, true // 반환값은 정의된 순서대로 반환한다
}

func MultiReturn2(a, b int) (result int, success bool) { // 반환값에 이름을 붙여서 반환한다.
	if b == 0 {
		result = 0
		success = false
		return
	}
	return a / b, true

}

func main() {
	c := Add(1, 2)         // Add 함수를 호출한다.
	fmt.Println(c)         // Add 함수의 반환값을 출력한다.
	fmt.Println(Add(1, 2)) // Add 함수를 호출한다.

	d, ok := MultiReturn(4, 2) // MultiReturn 함수를 호출한다.
	_ = ok
	_ = d
	fmt.Println(d, ok) // MultiReturn 함수의 반환값을 출력한다.

	f, succsecc := MultiReturn2(4, 2)                // MultiReturn2 함수를 호출한다.
	fmt.Printf("f: %d, succsecc: %t\n", f, succsecc) // MultiReturn2 함수의 반환값을 출력한다.

	fmt.Println(Fibo(2))

	g := Scaler(1, 2, Level1)
	h := Scaler(1, 2, Level2)
	i := Scaler(1, 2, Level3)
	fmt.Println("g:", g, "h:", h, "i:", i)
}

// #임수
