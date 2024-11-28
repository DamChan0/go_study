package main

import "fmt"

func Range(array []int) {
	for i, v := range array {
		fmt.Printf("array[%d] = %d\n", i, v)
	}

	println("just index")

	for _, v := range array {
		println(v)
	}

}

func main() {

	var array [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} // 배열을 선언하고 초기화한다.

	for i := 0; i < 10; i++ {
		fmt.Printf("array[%d] = %d\n", i, array[i]) // 배열의 요소를 출력한다.
	}

	initWithDeclare := [5]string{"A", "B", "C", "D", "E"} // 배열을 선언하고 초기화한다.

	for i := 0; i < 5; i++ {
		fmt.Printf("initWithDeclare[%d] = %s\n", i, initWithDeclare[i]) // 배열의 요소를 출력한다.
	}

	rangeArray := [5]int{1: 10, 3: 30}

	for i := 0; i < 5; i++ {
		fmt.Printf("rangeArray[%d] = %d\n", i, rangeArray[i])
		// 할당하지 않은 부분은 0으로 초기화된다.
	}

	insertArray := [10]int{11, 222, 3333, 44444, 555555}
	Range(insertArray[:])

}
