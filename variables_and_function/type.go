package main

import "fmt"

func printValue2() {
	var a int32 = 360
	var b int8 = int8(a)
	fmt.Println("a:", a)
	fmt.Printf("a : %d\n", a)
	fmt.Printf("a:0x%b\n", a) // 0x는 16진수를 나타내는 접두사
	fmt.Println("b:", b)
	fmt.Printf("b: 0x%b\n", b) // 0x는 16진수를 나타내는 접두사
}
