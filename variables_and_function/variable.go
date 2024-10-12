package main

import "fmt"

var count int = 0

func printValue() {
	intVariable := 10

	fmt.Println("hello", intVariable)
	count++
	fmt.Printf("count 출력 : %d \n ", count)
	var check bool
	check = count == 0

	fmt.Printf("count = 0 ? => %t\n", check) // boolean 값은 %t를 사용해서 출력할 수 았다
}

func main() {
	printValue()
	printValue2()

}
